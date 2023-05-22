#!/bin/bash

docker stop sharecvweb
docker rm sharecvweb
docker rmi lab/sharecvweb:latest
docker build -t lab/sharecvweb:latest .

docker run -p 6002:3000 --name sharecvweb --restart=always --network=app --ip 172.28.0.22 -d lab/sharecvweb:latest
