#bin/bash

export GO111MODULE=on
env GOOS=linux go build -ldflags="-s -w" -o bin/v1 v1/*.go