#!/bin/bash

echo "Authenticating to DuPontPublic"

# Change the password
DUPONT_GUEST_PASSWORD="M3rch"
curl -X POST \
  -d "f_user=guest&f_pass=${DUPONT_GUEST_PASSWORD}&submit=Log+In" \
  -H "Content-Type:application/x-www-form-urlencoded" \
  http://172.28.128.1:880/cgi-bin/hslogin.cgi

echo "Done!"
