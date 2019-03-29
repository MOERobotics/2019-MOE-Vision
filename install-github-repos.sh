#!/bin/bash

echo "Installing GitHub repos"

cd ~

git clone https://github.com/KHS-Robotics/v4l2-ctl-rest-api
git clone https://github.com/Ernie3/pi_h264.git

# PI H264
cd ~/pi_h264
npm install

# V4L2 REST API
cd ~/v4l2-ctl-rest-api
npm install

# Return home
cd ~

echo "Done!"
