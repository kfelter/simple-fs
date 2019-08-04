#!/bin/bash
EC2ADDR=18.234.196.95
echo
echo build the go bin
GOOS=linux go build -o portfolio.linux
echo attempt to deploy index.html to ec2 at $EC2ADDR
scp portfolio.linux ubuntu@$EC2ADDR:/home/ubuntu/simple-fs/portfolio.linux
