#!/bin/bash

echo $PWD
echo $(ls)

##build Go API
go get -d ../api all
go build ../api/main.go

##build Vue app
npm run build npm run build --prefix ../imperial/presentation/