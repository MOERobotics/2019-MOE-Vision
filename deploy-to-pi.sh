#!/bin/bash

echo "Deploy to Pi"

ssh -t pi@raspberrypi.local "mkdir -p /home/pi/go-vision"
scp -r ../* pi@raspberrypi.local:/home/pi/go-vision/

echo "Done!"
