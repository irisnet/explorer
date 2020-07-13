# explorer
Blockchain explorer for the IRIS Hub
##### 1 Prerequisites

* go1.13.0+

##### 2 install

```bash
    go get github.com/irisnet/explorer/backend
    make all
```

##### 3 Setting environment variables

```
    DB_ADDR          : mongo server's address
    DB_DATABASE      : database's name
    DB_USER          : database's username
    DB_PASSWORD      : database's password
    DB_POOL_LIMIT    : database max connection num

    PORT             : explorer server's port
    ADDR_NODE_SERVER : irishub lcd address
    ADDR_HUB_RPC     : irishub full node address
    FAUCET_URL       : faucet address
    CHAIN_ID         : irishub chain-id
    API_VERSION      : explorer api version
    MAX_DRAW_CNT     : Maximum number of collections
    SHOW_FAUCET      : switch of show faucet
    INITIAL_SUPPLY   : initial supplay of IRIS Token
    CUR_ENV          : current environment(dev/qa/testnet/mainnet)
    CronTimeAssetGateways: time interval of update asset gateways
    CronTimeAssetTokens: time interval of update asset tokens
    CronTimeGovParams: time interval of update gov params
    CronTimeTxNumByDay: time interval of update tx num by day
    CronTimeControlTask: time interval of monitor task execute
    CronTimeAccountRewards: time interval of update account rewards
    CronTimeValidators: time interval of update validators
    CronTimeValidatorIcons: time interval of update validator icons
    CronTimeProposalVoters: time interval of update voter info of proposal
    CronTimeValidatorStaticInfo: time interval of cronjob to update validator static info include uptime, selfBond, delegatorNum
    CronTimeFormatStaticDay: define time format of cronjob execute by every day
    CronTimeFormatStaticMonth: define time format of cronjob execute by eveny month
    CronTimeStaticDataDay: time interval of cronjob to snapshot delegator and validator rewards info
    CronTimeStaticDataMonth: time interval of cronjob to caculate delegator and validator rewards info
    CronTimeHeartBeat: time interval of heart beat in cron task
    NetreqLimitMax: max network request to lcd node
    
    
    //irishub v0.11.0 add
    PrefixAccAddr    : faa
    PrefixAccPub     : fap
    PrefixValAddr    : fva
    PrefixValPub     : fvp
    PrefixConsAddr   : fca
    PrefixConsPub    : fcp
    

```


##### 4 setup a light client for the IRIS Hub
After the irisplorer rest-server command is executed, a synchronization process will be started in the background, and the specified block chain transaction, account, block and other information will be synchronized to the local mongodb.
In addition, the command also opens a lightweight client program that provides a set of rest interfaces for the presentation of the irisplorer front view.

```golang
   irisplorer rest-server --mgo-url=localhost:27017 --chain-id=pangu --node=tcp://localhost:46657
```

##### 5 update the swagger file
At first, update the file in the docs dirname with  the command:
```golang
  #swag init
```

Then, copy the swagger.yaml content of docs to the lcd/lit/swagger-ui/swagger.yaml and cover  after the "swagger: \"2.0\"" position.
that's sure ok,than execute the command:
```golang
   #cd ackend/lcd/lite
   #go generate
```