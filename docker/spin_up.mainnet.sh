#!/bin/sh
docker-compose -f /Users/ivaylopetrov/Dev/janus/docker/quick_start/docker-compose.mainnet.yml up -d --build
# sleep 3 #executing too fast causes some errors
# docker cp ${GOPATH}/src/github.com/qtumproject/janus/docker/fill_user_account.sh qtum_testchain:.
# docker exec qtum_mainnet /bin/sh -c ./fill_user_account.sh