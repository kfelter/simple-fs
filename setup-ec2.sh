#!/bin/bash
echo setup the service to run continuously
sudo mv portfolio.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl start portfolio
sudo systemctl enable portfolio
sudo systemctl status portfolio