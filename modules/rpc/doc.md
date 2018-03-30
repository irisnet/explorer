# irisplorer rest api
The following describes a set of API interfaces for browser display data.

* Check the individual account balance
```
    /account/{address}
```
* Paged query bulk account balance.
```
    /accounts/{page}
```
* Paging query bulk block information.
```
    /blocks/{page}
```
* Query the latest block information.
```
    /blocks/recent
```
* Query a collection of validations for a block height.
```
    /validators/{height}
```
* Query for a block information.
```
    /block/{height}
```
* Query a collection of nodes for an account agent.
```
    /stake/candidates/{address}
```
* Enquiry of transaction details
```
    /tx/{hash}
```
* Paging query transfer transaction records.
```
    /tx/coin/{page}
```
* Check all transaction records of an account.
```
    /tx/coin/{address}
```
* Paging to query transaction records of an account.
```
    /tx/coin/{address}/{page}
```
* Paged query proxy class transaction records.
```
    /tx/stake/{page}
```
* Query all proxy transaction records for an account.
```
    /tx/stake/{address}
```
* Paging to query a proxy transaction record for an account.
```
    /tx/stake/{address}/{page}
```