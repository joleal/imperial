#!/bin/bash

echo 'build'
##build Go API
cd /usr/lib/imperial/api && go get all
sudo go build /usr/lib/imperial/api/main.go

##build Vue app
echo 'instal presentation packages'
cd /usr/lib/imperial/presentation && sudo npm install --unsafe-perm=true --allow-root

echo 'build presentation'
cd /usr/lib/imperial/presentation && sudo npm run build