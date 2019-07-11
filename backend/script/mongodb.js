//create database and user
// use sync-iris
// db.createUser(
//     {
//         user:"iris",
//         pwd:"irispassword",
//         roles:[{role:"root",db:"admin"}]
//     }
// )

db.createCollection("ex_power_change");
db.createCollection("ex_uptime_change");
db.createCollection("ex_tx_num_stat");
db.createCollection("ex_config");
db.createCollection("ex_val_black_list");
db.createCollection("ex_validator");
db.createCollection("ex_gov_params");

// create index
db.ex_power_change.createIndex({"height": 1, "address": 1}, {"unique": true});
db.ex_tx_num_stat.createIndex({"date": -1}, {"unique": true});
db.ex_config.createIndex({"env_nm": 1, "chain_id": 1}, {"unique": true, "background": true});
db.ex_val_black_list.createIndex({"operator_addr": 1}, {"unique": true});
db.ex_validator.createIndex({"operator_address": 1}, {"unique": true});
db.ex_validator.createIndex({"proposer_addr": 1}, {"unique": true,"background": true});
db.ex_gov_params.createIndex({"key": 1}, {"unique": true});

// init data
db.ex_config.insert({"env_nm":"mainnet","host":"https://www.irisplorer.io","chain_id":"irishub","show_faucet":0});
db.ex_config.insert({"env_nm":"testnet","host":"https://testnet.irisplorer.io","chain_id":"fuxi","show_faucet":1});
db.ex_config.insert({
    "env_nm": "testnet",
    "host": "http://nyancat.irisplorer.io",
    "chain_id": "nyancat",
    "show_faucet": 1
});
db.ex_config.insert({"env_nm": "dev", "host": "http://dev.irisplorer.io", "chain_id": "rainbow-dev", "show_faucet": 1});
db.ex_config.insert({"env_nm": "qa", "host": "http://qa.irisplorer.io", "chain_id": "rainbow-qa", "show_faucet": 1});

db.ex_val_black_list.insert({
    "operator_addr": "iva18claj4r9x3gj5yurjxec29p2c9x6t49r6dqp00",
    "moniker": "Validator20190320-1",
    "identity": "",
    "website": "",
    "details": ""
});
