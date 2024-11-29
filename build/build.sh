#!/bin/bash

shot_ID=$(git rev-parse --short HEAD)

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'main.Version=${shot_ID}'" -o output/cheetach-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X 'main.Version=${shot_ID}'" -o output/cheetach-linux-arm64

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X 'main.Version=${shot_ID}'" -o output/cheetach-windows-amd64

CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-X 'main.Version=${shot_ID}'" -o output/cheetach-darwin-amd64
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -ldflags "-X 'main.Version=${shot_ID}'" -o output/cheetach-darwin-arm64