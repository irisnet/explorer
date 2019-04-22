//create database and user
// use sync-iris
// db.createUser(
//     {
//         user:"iris",
//         pwd:"irispassword",
//         roles:[{role:"root",db:"admin"}]
//     }
// )

db.createCollection("block");
db.createCollection("stake_role_candidate");
db.createCollection("sync_task");
db.createCollection("tx_common");
db.createCollection("proposal");
db.createCollection("tx_msg");
db.createCollection("power_change");
db.createCollection("uptime_change");
db.createCollection("sync_conf");
db.createCollection("mgo_txn");
db.createCollection("mgo_txn.stash");
db.createCollection("ex_tx_num_stat");
db.createCollection("ex_config");
db.createCollection("val_black_list");
db.createCollection("validator");

// create index
db.block.createIndex({"height": -1}, {"unique": true});

db.stake_role_candidate.createIndex({"address": 1}, {"unique": true});
db.stake_role_candidate.createIndex({"pub_key": 1});

db.sync_task.createIndex({"start_height": 1, "end_height": 1}, {"unique": true});

db.tx_common.createIndex({"height": -1});
db.tx_common.createIndex({"time": -1});
db.tx_common.createIndex({"tx_hash": 1}, {"unique": true});
db.tx_common.createIndex({"from": 1});
db.tx_common.createIndex({"to": 1});
db.tx_common.createIndex({"type": 1});
db.tx_common.createIndex({"status": 1});

db.proposal.createIndex({"proposal_id": 1}, {"unique": true});

db.tx_msg.createIndex({"hash": 1}, {"unique": true});

db.power_change.createIndex({"height": 1, "address": 1}, {"unique": true});

db.ex_tx_num_stat.createIndex({"date": -1}, {"unique": true});

db.ex_config.createIndex({"env_nm": 1}, {"unique": true});

db.val_black_list.createIndex({"operator_addr": 1}, {"unique": true});

db.validator.createIndex({"operator_address": 1}, {"unique": true});


// init data
db.sync_conf.insert({"block_num_per_worker_handle": 100, "max_worker_sleep_time": 120});

db.ex_config.insert({"env_nm":"mainnet","host":"https://www.irisplorer.io","chain_id":"irishub","show_faucet":0});
db.ex_config.insert({"env_nm":"testnet","host":"https://testnet.irisplorer.io","chain_id":"fuxi","show_faucet":1});
db.ex_config.insert({"env_nm":"NYAN_CATS","host":"https://nyancat.bigdipper.live/","chain_id":"nyan_cats","show_faucet":0});

db.val_black_list.insert({"operator_addr":"iva18claj4r9x3gj5yurjxec29p2c9x6t49r6dqp00","moniker":"Validator20190320-1","identity":"","website":"","details":""});

