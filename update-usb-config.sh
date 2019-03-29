getVideoCameraStreamPath(){
  v4l2-ctl --list-devices | awk "/$1/{getline; print}" | grep -E -o "/dev.*"
}

microsoftPath=$(getVideoCameraStreamPath "Microsoft")
logitechPath=$(getVideoCameraStreamPath "C920")

echo ${microsoftPath}
echo ${logitechPath}

microsoftJson=$(cat ~/pi_h264/config.json | jq ".device = \"$microsoftPath\"")
echo $microsoftJson > ~/pi_h264/config.json

logitechNumber=$(echo ${logitechPath} | grep -E -o "[0-9]*")

v4l2-ctl --device=${logitechNumber} --set-ctrl=exposure_auto=1
v4l2-ctl --device=${logitechNumber} --set-ctrl=exposure_absolute=80
v4l2-ctl --device=${logitechNumber} --set-ctrl=contrast=255

sudo systemctl daemon-reload

sudo systemctl stop h264server
sudo systemctl stop go-vision

sudo systemctl restart h264server
sudo systemctl restart go-vision
