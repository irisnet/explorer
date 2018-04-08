#!/usr/bin/env bash
####build irisplorer
cd ../server/
make irisplorer

### make irisplorer image
docker build -t irisplorer:0.1.0 -f ./Dockerfile .

### make irisplorer ui image
cd ../ui/
yarn install
docker build -t irisplorer-ui:0.1.0 -f ./Dockerfile .

cd ../setup/
docker network create irisplorernet
docker-compose -f ./docker-compose.yml up -d
