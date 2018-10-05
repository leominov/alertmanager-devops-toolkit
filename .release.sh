#!/bin/bash

set -x

VERSION=$(cat main.go | grep "Version =" | awk -F\" '{print $2}')
FILENAME="alertmanager-devops-toolkit-$VERSION.linux-amd64"

GOOS=linux go build -o "$FILENAME"
dope nexus upload -d prometheus "$FILENAME"
rm -f "$FILENAME"
