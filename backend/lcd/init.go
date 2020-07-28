package lcd

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/weichang-bianjie/irishub-sdk-go"
	"github.com/weichang-bianjie/irishub-sdk-go/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"time"
	"path/filepath"
	"fmt"
)

var (
	client sdk.Client
)

func init() {
	path := filepath.Join("/User/user/go/", "test")
	client = sdk.NewClient(types.ClientConfig{
		NodeURI: conf.Get().Hub.NodeUrl,
		Network: types.Testnet,
		ChainID: conf.Get().Hub.ChainId,
		Mode:    types.Commit,
		Timeout: time.Duration(10 * time.Second),
		Fee: []types.DecCoin{
			{Denom: "stake", Amount: types.NewDecimal(1)},
		},
		Gas:       2000000,
		KeyDAO:    types.NewMemoryDB(), //default keybase
		StoreType: types.PrivKey,
		Level:     "info",
		DBRootDir: path,
	})


	govParamMap, err := GetGovModuleParamMap(GovModule)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if v, ok := govParamMap["tally_params"]; ok {
		if data, ok := v.(map[string]interface{}); ok {
			if threshold, ok := data[NormalThresholdKey].(string); ok {
				NormalThreshold = threshold
			}
			if veto, ok := data[NormalVetoKey].(string); ok {
				NormalVeto = veto
			}
			if quorum, ok := data[NormalQuorumKey].(string); ok {
				NormalParticipation = quorum
			}

		}
	}

	if v, ok := govParamMap["deposit_params"]; ok {
		if data, ok := v.(map[string]interface{}); ok {
			depositparams, ok := data[NormalMinDepositKey].([]interface{})
			if ok {

				if len(depositparams) > 0 {
					first := depositparams[0].(map[string]interface{})
					coinAsUtils := utils.ParseCoin(fmt.Sprintf("%v%v", first["amount"], first["denom"]))
					NormalMinDeposit.Amount = coinAsUtils.Amount
					NormalMinDeposit.Denom = coinAsUtils.Denom
				}
			}
		}
	}

}