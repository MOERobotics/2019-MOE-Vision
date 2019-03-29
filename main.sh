#!/bin/bash

echo "Install service"

./build-go-vision.sh
./install-service.sh

echo "Done!"
