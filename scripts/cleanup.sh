#!/bin/bash
echo 'clean'
echo 'remove main executable'
sudo rm -rf /usr/lib/imperial/api/main

echo 'remove presentation dist'
sudo rm -rf /usr/lib/imperial/presentation/dist
