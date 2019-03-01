//create database and user
// use sync-iris
// db.createUser(
//     {
//         user:"iris",
//         pwd:"irispassword",
//         roles:[{role:"root",db:"admin"}]
//     }
// )


db.createCollection("ex_config");

// create index
db.ex_config.createIndex({"env_nm": 1}, {"unique": true});

// init data
db.ex_config.insert({"env_nm":"mainnet","host":"https://www.irisplorer.io","chain_id":"fuxi","show_faucet":0});
db.ex_config.insert({"env_nm":"testnet","host":"https://testnet.irisplorer.io","chain_id":"fuxi","show_faucet":1});
db.ex_config.insert({"env_nm":"NYAN_CATS","host":"https://nyancat.bigdipper.live/","chain_id":"fuxi","show_faucet":0});

