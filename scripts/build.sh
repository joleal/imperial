#!/bin/bash

echo 'build'
##build Go API
cd /usr/lib/imperial/api && go get all
sudo go build /usr/lib/imperial/api/main.go

##build Vue app
echo 'instal presentation packages'
sudo npm install --prefix /usr/lib/imperial/presentation/

echo 'build presentation'
sudo npm run build npm run build --prefix /usr/lib/imperial/presentation/