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
db.createCollection("val_black_list");
db.createCollection("validator");

// create index
db.ex_config.createIndex({"env_nm": 1}, {"unique": true});
db.val_black_list.createIndex({"operator_addr": 1}, {"unique": true});
db.validator.createIndex({"operator_addr": 1}, {"unique": true});

// init data
db.ex_config.insert({"env_nm":"mainnet","host":"https://www.irisplorer.io","chain_id":"irishub","show_faucet":0});
db.ex_config.insert({"env_nm":"testnet","host":"https://testnet.irisplorer.io","chain_id":"fuxi","show_faucet":1});
db.ex_config.insert({"env_nm":"NYAN_CATS","host":"https://nyancat.bigdipper.live/","chain_id":"nyan_cats","show_faucet":0});
db.val_black_list.insert({"operator_addr":"iva18claj4r9x3gj5yurjxec29p2c9x6t49r6dqp00","moniker":"Validator20190320-1","identity":"","website":"","details":""});

