#!/bin/bash

echo 'start'
echo $(ls /usr/lib/imperial)

##build Go API
sudo go get -d /usr/lib/imperial/api all
sudo go build /usr/lib/imperial/api/main.go

##build Vue app
sudo npm run build npm run build --prefix /usr/lib/imperial/presentation/