#!/usr/bin/env bash
buildServer(){
    cd ../server/
    if [ ! -f "./build/irisplorer" ];then
    echo "make irisplorer"
    make irisplorer
    else
    echo "irisplorer has existed,skip [make irisplorer]"
    fi

    echo "build docker images[irisplorer:0.1.0]"
    docker build -t irisplorer:0.1.0 -f ./Dockerfile .
}

buildUi(){
    cd ../ui/
    yarn install
    echo "build docker images[irisplorer-ui:0.1.0]"
    docker build -t irisplorer-ui:0.1.0 -f ./Dockerfile .
}

up(){
    cd ../setup/
    docker network create irisplorernet
    echo "start docker container..."
    docker-compose -f ./docker-compose.yml up -d
}

down(){
    cd ../setup/
    echo "stop docker container ..."
    docker-compose -f ./docker-compose.yml down
}

execute(){
    if [ $1 = "server" ];then
        buildServer
    elif [ $1 = "ui" ]; then
        buildUi
    elif [ $1 = "u" ]; then
        up
    elif [ $1 = "d" ]; then
        down
    elif [ $1 = "all" ]; then
        buildServer
        buildUi
        up
    fi
}

echo "exec $1"
execute $1