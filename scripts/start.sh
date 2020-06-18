#!/bin/bash
echo 'start'
sudo service imperial restart
sudo nginx -s reload