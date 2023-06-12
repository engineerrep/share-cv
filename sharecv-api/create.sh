#!/bin/bash

if [ ! -f go.mod ];then go mod init sharecvapi;fi
if [ ! -d deploy ];then mkdir deploy;fi

if [ ! -d deploy/etc ];then mkdir deploy/etc;fi
if [ ! -d deploy/logs ];then mkdir deploy/logs;fi
if [ ! -d deploy/key ];then mkdir deploy/key;fi
if [ ! -d deploy/mmdb ];then mkdir deploy/mmdb;fi

if [ ! -d deploy/doc ];then mkdir deploy/doc;fi
if [ ! -d deploy/doc/api ];then mkdir deploy/doc/api;fi

base_dir=$(pwd)

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy

go install github.com/zeromicro/go-zero/tools/goctl@latest
go install github.com/zeromicro/goctl-swagger@latest
