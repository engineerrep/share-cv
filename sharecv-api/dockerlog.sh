#!/bin/bash

docker logs -f -t --since="2023" --tail=10 sharecvapi
