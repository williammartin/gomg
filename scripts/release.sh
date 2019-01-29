#!/usr/bin/env bash

workspace=$(mktemp -d)

for platform in linux darwin; do
  echo "building ${platform}..."
  env GOOS=${platform} GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ${workspace}/gomg_${platform} github.com/williammartin/gomg
done

open ${workspace}
