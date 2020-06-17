#!/bin/bash

##build Go API
go get -d /home/ubuntu/imperial/api all
go build /home/ubuntu/imperial/api/main.go

##build Vue app
npm run build npm run build --prefix /home/ubuntu/imperial/presentation/