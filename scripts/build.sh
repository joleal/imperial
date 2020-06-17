#!/bin/bash

echo 'start'
echo $(ls /usr/lib/imperial)

##build Go API
cd /usr/lib/imperial/api && go get all
sudo go build /usr/lib/imperial/api/main.go

##build Vue app
sudo npm run build npm run build --prefix /usr/lib/imperial/presentation/