# explorer
Blockchain explorer for the IRIS Hub
##### 1 Prerequisites

* go1.9.2+

##### 2 install

```golang
    go get github.com/irisnet/explorer
    make all
```

##### 3 setup a light client for the IRIS Hub
After the irisplorer rest-server command is executed, a synchronization process will be started in the background, and the specified block chain transaction, account, block and other information will be synchronized to the local mongodb.
In addition, the command also opens a lightweight client program that provides a set of rest interfaces for the presentation of the irisplorer front view.

```golang
   irisplorer rest-server --mgo-url=localhost:27017 --chain-id=pangu --node=tcp://localhost:46657
```

##### 4 api doc for developer 
- See the [Rest Guide](rest/controller/doc.md)