#!/bin/bash

# usge: build/build.sh release or build/build.sh debug

set -x

export SERVICE_NAME="smartplug"
export GOPROXY="https://goproxy.cn"

TAG="latest"

function LogOut()
{
	echo "`date "+%Y-%m-%d %H:%M:%S"` " $@
}

function docker_check()
{
    v=`docker --version`
    if [[ $? != 0 ]]; then
        LogOut "Need to install docker first"
        exit 1
    fi
    LogOut "docker_check ok, $v"
}

function build()
{
    err=`go build -o ${SERVICE_NAME}`
    if [[ $? != 0 ]]; then
        LogOut "build failed, error: " $err
        exit 1
    fi
    ls -l |grep $SERVICE_NAME
    LogOut "build ok"
}

function docker_build()
{
    pushd build
    cp -r ../conf .
    cp -r ../static .
    cp ../${SERVICE_NAME} .
    chmod +x ${SERVICE_NAME}
    chmod +x entrypoint.sh
    
    info=`sudo docker build -t ${SERVICE_NAME}:${TAG} .`
    if [[ $? != 0 ]]; then
        LogOut "docker build failed, error: " $info
        exit 1
    fi
    LogOut "docker build ok"

    rm -rf conf
    rm -rf static
    rm -rf ${SERVICE_NAME}

    popd
}

function docker_save()
{
    pushd build
    info=`sudo docker save -o ${SERVICE_NAME}_${TAG}.tar ${SERVICE_NAME}:${TAG}`
    if [[ $? != 0 ]]; then
        LogOut "docker save failed, error: " $info
        exit 1
    fi
    LogOut "docker save ok, $SERVICE_NAME_$TAG.tar"
    popd
}

if [[ $# = 1 ]]; then
    if [[ $1 = "release" ]]; then
        TAG=`date "+%Y%m%d%H%M%S"`
    elif [[ $1 = "debug" ]]; then
        TAG=`date "+%Y%m%d%H%M%S"`
    fi
fi

docker_check
build
docker_build
docker_save
