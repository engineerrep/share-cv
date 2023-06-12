#!/bin/bash

git checkout .
git pull

chmod +x create.sh
sh create.sh

BASE_PATH=$(pwd)

# ****************** Build service and doc *****************************

# api api service
cd "$BASE_PATH"/app/api && make release

# 构建镜像
cd "$BASE_PATH" && docker build --no-cache -t lab/sharecvapi:latest -f Dockerfile .

# 启动镜像
docker-compose down
docker-compose up -d

docker rmi -f $(docker images | grep "none" | awk '{print $3}')
