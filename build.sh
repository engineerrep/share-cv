#!/bin/bash

rm -rf node_modules .output .nuxt yarn.lock package-lock.json

node -v

npm -v

npm install

npm run build:pre
