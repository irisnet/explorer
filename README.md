# Explorer

Blockchain explorer for IRIS Hub,which includes explorer-backend and explorer-frontend.

# Explorer-backend

## Run with docker
You can run application with docker.

```$xslt
cd github.com/irisnet/explorer/backend
docker build -t explorer-backend .
```
## ENV

DB_ADDR: mongo url ( eg. 192.168.150.7:27017)

DB_DATABASE: mongo database name ( eg. sync-iris )

DB_USER: mongo user ( eg. iris )

DB_PASSWORD: mongo password ( eg. irispassword )

FAUCET_URL: faucet url ( eg. http://dev.faucet.irisplorer.io )

CHAIN_ID: chain-id ( eg. rainbow-dev )

detail: github.com/irisnet/explorer/backend/README.md

## Docker Run
docker run -p 8080:8080 -e ${ENV Variables} explorer-backend:latest

# Explorer-frontend

##Config
- vue.config.js
```$xslt
module.exports = {
    devServer: {
        proxy: {
          '/api':{
              target:'http://localhost:8080',        //backend address:port
              ......
          }
        }
    },
};
```


## Run with docker
You can run application with docker.

```$xslt
cd github.com/irisnet/explorer/frontend
docker build -t explorer-frontend .
```

## Docker Run
docker run -p 8080:8080 -e ${ENV Variables} explorer-frontend:latest


