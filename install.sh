#!/bin/bash
#
# this script downloads the newest version from github.com and installs it to /usr/local/bin
#

GITLOCATION="https://raw.githubusercontent.com/hunterlong/storj-go/master/storj-go"

printf "\n storj-go - Storj.io API Wrapper\n"
printf " install location: '/usr/local/bin/storj-go'\n"

sudo curl $GITLOCATION > /usr/local/bin/storj-go
sudo chmod +x /usr/local/bin/storj-go

printf "Run it with: 'storj-go'\n"
