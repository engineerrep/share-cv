#!/bin/bash

chmod +x create.sh
sh create.sh

BASE_PATH=$(pwd)

# ****************** Build service and doc *****************************

# app api service
cd "$BASE_PATH"/app/api && make all
