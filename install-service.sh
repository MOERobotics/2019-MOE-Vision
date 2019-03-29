#!/bin/bash

echo "Install service"

GO_VISION_DIR=/home/pi/go-vision

sudo cp ${GO_VISION_DIR}/resources/go-vision.service \
  /etc/systemd/system/go-vision.service

sudo systemctl daemon-reload
sudo systemctl enable go-vision.service
sudo systemctl start go-vision.service
sudo systemctl restart go-vision.service

echo "Done!"
