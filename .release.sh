#!/bin/bash

set -x

GOOS=linux
VERSION=$(cat main.go | grep "Version =" | awk -F\" '{print $2}')
FILENAME="alertmanager-devops-toolkit-$VERSION.linux-amd64"

go build -o "$FILENAME"
dope nexus upload -d prometheus "$FILENAME"
rm -f "$FILENAME"
