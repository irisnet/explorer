package task

import (
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var validatorMap = make(map[string]types.PowerChange)

type ResValidatorPreCommits struct {
	Id            ID    `bson:"_id"`
	PreCommitsNum int64 `bson:"num"`
}

type ID struct {
	Address string `bson:"address"`
	Hour    string `bson:"hour"`
}

func Start() {
	for true {
		UptimeChange()
	}
}

func init() {
	var powerChanges []types.PowerChangeOrder
	db := utils.GetDatabase()
	p := db.C("power_change")
	pipe := p.Pipe(
		[]bson.M{
			{"$group": bson.M{
				"_id":    "$address",
				"power":  bson.M{"$last": "$power"},
				"date":   bson.M{"$last": "$date"},
				"change": bson.M{"$last": "$change"},
			}},
		},
	)
	pipe.All(&powerChanges)

	for _, powerChange := range powerChanges {
		if powerChange.Change != types.Slash {
			validatorMap[powerChange.Address] = types.PowerChange{Address: powerChange.Address, Time: powerChange.Time, Power: powerChange.Power, Change: powerChange.Change}
		}
	}
}

func UptimeChange() {
	log.Println("task start")
	var uptimeChange types.UptimeChange
	var blocks []document.Block
	var firstBlock document.Block
	var lastBlock document.Block

	db := utils.GetDatabase()
	b := db.C("block")
	u := db.C("uptime_change")
	p := db.C("power_change")

	defer db.Session.Close()

	u.Find(nil).Sort("-time").One(&uptimeChange)

	var err error
	var startTime time.Time
	d, _ := time.ParseDuration("1h") //1 hour later
	if uptimeChange.Time == "" {
		b.Find(nil).Sort("height").One(&firstBlock)
		startTime = firstBlock.Time
	} else {
		startTime, _ = time.ParseInLocation("2006-01-02 15", uptimeChange.Time, time.UTC)
		startTime = startTime.Add(d)
	}

	err = b.Find(nil).Sort("-height").One(&lastBlock)
	if err != nil {
		log.Println(err.Error())
		return
	}

	lastTime := lastBlock.Time
	endTime := startTime.Add(d)

	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 0, 0, 0, startTime.Location())
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), 0, 0, 0, endTime.Location())

	log.Printf("startTime:%s,endTime:%s", startTime.UTC().Format("2006-01-02 15"), endTime.UTC().Format("2006-01-02 15"))

	if !endTime.Before(lastTime) {
		log.Printf("handle to now %s, task end\n", startTime.Format("2006-01-02 15"))
		time.Sleep(10 * time.Minute)
		return
	}

	b.Find(bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}}).Sort("height").All(&blocks)

	if err != nil {
		log.Println(err.Error())
		return
	}

	//init validatorMap if validatorMap length is 0
	if len(validatorMap) == 0 && len(blocks) > 0 {
		for _, validator := range blocks[0].Validators {
			powerChange := types.PowerChange{
				Address: validator.Address,
				Power:   validator.VotingPower,
				Time:    blocks[0].Time,
				Change:  types.Change,
			}
			validatorMap[validator.Address] = powerChange
			p.Insert(powerChange)
		}
	}

	uptimeMap := make(map[string]int)
	for _, block := range blocks {

		// power change handle
		validatorMapNow := make(map[string]int64)
		for _, validator := range block.Validators {
			validatorMapNow[validator.Address] = validator.VotingPower

			// validator add or update
			if _, ok := validatorMap[validator.Address]; !ok {
				powerChange := types.PowerChange{
					Address: validator.Address,
					Power:   validator.VotingPower,
					Time:    block.Time,
					Height:  block.Height,
					Change:  types.Change,
				}
				validatorMap[validator.Address] = powerChange
				p.Insert(&powerChange)
			} else {
				if validatorMap[validator.Address].Power != validator.VotingPower {
					powerChange := types.PowerChange{
						Address: validator.Address,
						Power:   validator.VotingPower,
						Time:    block.Time,
						Height:  block.Height,
						Change:  types.Change,
					}
					validatorMap[validator.Address] = powerChange
					p.Insert(&powerChange)
				}
			}
		}

		// validator slash
		for k, v := range validatorMap {
			if v.Address == "" {
				continue
			}
			if _, ok := validatorMapNow[k]; !ok {
				powerChange := types.PowerChange{
					Address: v.Address,
					Power:   v.Power,
					Time:    block.Time,
					Height:  block.Height,
					Change:  types.Slash,
				}
				delete(validatorMap, k)
				p.Insert(&powerChange)
			}
		}

		// uptime handle
		for _, validator := range block.Validators {
			if _, ok := uptimeMap[validator.Address]; !ok {
				uptimeMap[validator.Address] = 0
			}
		}
		for _, commit := range block.Block.LastCommit.Precommits {
			uptimeMap[commit.ValidatorAddress]++
		}
	}

	doneTime := startTime.UTC().Format("2006-01-02 15")
	for k, v := range uptimeMap {
		uptimeChange := types.UptimeChange{Address: k, Uptime: float64(100*v) / float64(len(blocks)), Time: doneTime}
		u.Insert(&uptimeChange)
	}
	log.Println("task end")
}
