#!/usr/bin/env node

let fs = require('fs');
let program = require('commander');

function list (val) {
    return val.split(',')
}

console.log('Replacing environments ...');
let accAddr = "fva", accPub = "fap", valAddr = "fva";



replaceEnv([
    "node_modules/irisnet-crypto/config.json"
    ],
    {
        "accAddr": accAddr,
        "accPub":accPub,
        "valAddr":valAddr
    }
);


function replaceEnv(files, params) {
    files.forEach(function(file,index){
        let result = fs.readFileSync(file).toString();
        for (let key in params) {
            let r = "\\${"+key+"}";
            result = result.replace(new RegExp(r,"g"), params[key]);
        }

        fs.writeFileSync(file, result)
    });
}
