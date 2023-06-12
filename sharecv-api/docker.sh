#!/bin/bash

docker stop sharecvapi
docker rm sharecvapi
docker rmi lab/sharecvapi:latest
docker build -t lab/sharecvapi:latest .

docker run -p 6000:6000 --name sharecvapi -v /data/sharecvapi/deploy:/app --restart=always --network=app --ip 172.28.0.20 -d lab/sharecvapi:latest

# 不自动启动 请手动执行启动
#docker stop sharecvapidoc
#docker rm sharecvapidoc
#docker run -p 6001:8080 --name sharecvapidoc -v /data/sharecvapi/deploy/doc/api:/foo -e SWAGGER_JSON=/foo/core.json --restart=always --network=app --ip 172.28.0.21 -d swaggerapi/swagger-ui:v4.15.5
