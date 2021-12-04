#!/bin/bash

BRANCH=${BRANCH:?'missing BRANCH env var'}
GOARCH=${GOARCH:-$(go env GOARCH)}

mkdir ./out

#go build -ldflags "-X 'main.Version=${BRANCH}(${GIT_SHA})' -X 'main.BuildTime=${DATE}" -o ./out/clymene-agent ./cmd/agent/main.go
CGO_ENABLED=0 go build -ldflags "-X 'main.Version=${BRANCH}'" -o ./out/clymene-agent ./cmd/agent/main.go

cp ./cmd/agent/Dockerfile ./

docker build -t bourbonkk/clymene-agent:"${BRANCH}" .
docker push bourbonkk/clymene-agent:"${BRANCH}"