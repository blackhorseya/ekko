#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v entity | grep -v docs); do
    go test -race -coverprofile=profile.out -covermode=atomic --tags=unit "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
