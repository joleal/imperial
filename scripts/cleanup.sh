#!/bin/bash
echo 'clean'
echo 'remove main executable'
sudo rm -rf /usr/lib/imperial/api/main

echo 'remove presentation dist'
sudo rm -rf /usr/lib/imperial/presentation/dist

echo 'run database scripts'
for f in `ls /usr/lib/imperial/scripts/database/*.sql | sort -V`
do
  echo "Processing $f file..."
  # take action on each file. $f store current file name
  mariadb imperial -u ubuntu < $f
done
