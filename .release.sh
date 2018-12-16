#!/bin/bash

set -e
set -x

VERSION=$(cat main.go | grep "Version =" | awk -F\" '{print $2}')
TARGETS=(
    "darwin"
    "linux"
)

for target in ${TARGETS[@]}; do
    filename="alertmanager-devops-toolkit-$VERSION.$target-amd64"
    GOOS="$target" go build -o ".build/$filename"
    github-release upload \
        --user "leominov" \
        --repo "alertmanager-devops-toolkit" \
        --tag "v$VERSION" \
        --name "$filename" \
        --file ".build/$filename"
done

for target in ${TARGETS[@]}; do
    dope nexus upload -d prometheus ".build/$filename"
done
