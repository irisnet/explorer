#!/usr/bin/env node

var fs = require('fs');
var program = require('commander');
var version = fs.readFileSync("version").toString().replace(/[\r\n]/g,"");
function list (val) {
  return val.split(',')
}
program
  .version(version)
  .option("-p, --params <items>", "An list of explorer environment and bamboo build number, e.g.: dev,25", list)
  .parse(process.argv);
console.log('Replacing environments ...');
var env = "dev";
var buildNum = "0";

if (program.params) {
  env = program.params[0] ? program.params[0] : env;
  buildNum = program.params[1] ? program.params[1] : buildNum;
}
replaceEnv("testVersion.js",
  {
    "version": version,
    "env": env,
    "buildNumber": buildNum,
  }
);


function replaceEnv(file, params) {
    var result = fs.readFileSync(file).toString();
    for (var key in params) {
      result = result.replace(new RegExp(`_${key}`,"i"), params[key]);
    }

    fs.writeFileSync(file, result)
}
