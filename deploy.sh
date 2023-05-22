#!/bin/bash

git checkout .
git pull

rm -rf node_modules .output .nuxt yarn.lock package-lock.json

npm install
npm run build:prod

chmod u+x docker.sh
sh docker.sh
