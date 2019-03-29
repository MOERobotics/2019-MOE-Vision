#!/bin/bash

echo "Build Go Vision"

GO_VISION_DIR=/home/pi/go-vision

go build -o ${GO_VISION_DIR}/bin/go-vision ${GO_VISION_DIR}/src/go-vision.go

echo "Done!"
