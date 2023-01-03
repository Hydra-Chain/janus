#!/bin/sh
docker-compose --env-file ./.env  -f ./quick_start/docker-compose.testnet.yml up -d --build
# sleep 3 #executing too fast causes some errors
# docker cp ${GOPATH}/src/github.com/qtumproject/janus/docker/fill_user_account.sh qtum_testchain:.
# docker exec qtum_mainnet /bin/sh -c ./fill_user_account.sh