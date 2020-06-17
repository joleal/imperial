#!/bin/bash

echo 'start'
echo $(ls /usr/lib/imperial)

##build Go API
go get -d ../api all
go build ../api/main.go

##build Vue app
npm run build npm run build --prefix ../imperial/presentation/