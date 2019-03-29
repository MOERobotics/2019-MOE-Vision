#!/bin/bash

echo "Executing Pi Package upgrade"

function go_vision::pi_upgrade() {
  sudo apt -y update &&
    sudo apt -y upgrade &&
    sudo apt -y dist-upgrade &&
    sudo apt -y full-upgrade
}

## Upgrade
go_vision::pi_upgrade

## Developer tools
sudo apt -y install \
  coreutils \
  curl \
  gcc \
  git \
  htop \
  jq \
  lsof \
  nano \
  nginx \
  nmap \
  p7zip \
  tree \
  vim \
  wget \
  network-manager

## Video utils
sudo apt -y install \
  v4l-utils

## FFMPEG

# sudo deb-src http://mirror.ox.ac.uk/sites/archive.raspbian.org/archive/raspbian/ stretch main contrib non-free rpi
# sudo deb-src http://archive.raspbian.org/raspbian/ stretch main contrib non-free rpi
# sudo deb-src http://www.deb-multimedia.org stretch main non-free

# sudo apt -y install \
#   build-essential \
#   dh-make \
#   fakeroot \
#   yasm \
#   pkg-config \
#   libfdk-aac-dev \
#   libx264-dev

sudo sed -i '$a deb http://www.deb-multimedia.org stretch main non-free' /etc/apt/sources.list

# TODO: Check if file already exists
wget http://www.deb-multimedia.org/pool/main/d/deb-multimedia-keyring/deb-multimedia-keyring_2016.8.1_all.deb
# TODO: Check if already installed
sudo dpkg -i deb-multimedia-keyring_2016.8.1_all.deb

go_vision::pi_upgrade && sudo apt -y install \
  ffmpeg

## Node
curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
sudo apt -y install \
  nodejs

## Go
GO_DOWNLOAD_FILE="go1.12.1.linux-armv6l.tar.gz"
# TODO: Check if file already exists
wget https://dl.google.com/go/${GO_DOWNLOAD_FILE}
# TODO: Check if location already has go installed
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xvf ${GO_DOWNLOAD_FILE}

## Go packages
go get -u all
go get -u gonum.org/v1/gonum/stat
go get -u gocv.io/x/gocv
go get -u github.com/hybridgroup/mjpeg

## Upgrade
go_vision::pi_upgrade

## Reboot
echo "Done! Now rebooting"
sleep 2
sudo reboot
