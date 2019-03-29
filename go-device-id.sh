#!/bin/sh

getVideoCameraStreamPath(){
  v4l2-ctl --list-devices | awk "/$1/{getline; print}" | grep -E -o "/dev.*"
}

logitechPath=$(getVideoCameraStreamPath "C920")

echo $(echo $logitechPath | grep -E -o "[0-9]*")

