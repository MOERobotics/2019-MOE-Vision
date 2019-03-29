#!/bin/bash

sudo nmcli connection modify "Wired connection 1" ipv4.never-default true

nmcli r wifi on
nmcli d wifi list
nmcli d wifi connect DuPontPublic
sudo nmcli d wifi connect DuPontPublic

sudo nmcli connection

sudo systemctl restart NetworkManager

sudo rm /run/wpa_supplicant/wlan0
