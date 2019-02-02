# explorer
Blockchain explorer for the IRIS network


# ENV
DB_ADDR: mongo url ( eg. 192.168.150.7:27017)

DB_DATABASE: mongo database name ( eg. sync-iris )

DB_USER: mongo user ( eg. iris )

DB_PASSWORD: mongo password ( eg. irispassword )

FAUCET_URL: faucet url ( eg. http://dev.faucet.irisplorer.io )

CHAIN_ID: chain-id ( eg. rainbow-dev )

# Docker Run
docker run -p 8080:8080 -e ${ENV Variables} explorer

