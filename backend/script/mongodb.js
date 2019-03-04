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