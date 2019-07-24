#!/bin/bash
EC2ADDR=18.234.196.95
echo
echo attempt to deploy index.html to ec2 at $EC2ADDR
scp static/index.html ubuntu@$EC2ADDR:/home/ubuntu/simple-fs/static/index.html
