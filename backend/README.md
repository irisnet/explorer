# explorer
Blockchain explorer for the IRIS Hub
##### 1 Prerequisites

* go1.9.2+

##### 2 install

```bash
    go get github.com/irisnet/explorer
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