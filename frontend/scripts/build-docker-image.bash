#!/usr/bin/env bash

VERSION=$(git describe --tags | cut -c 2-)
echo "${VERSION}" | figlet

docker buildx build \
  --platform linux/amd64 \
  --push \
  --tag "registry.zat.ong/bahn-alarm-frontend:${VERSION}" \
  --tag "registry.zat.ong/bahn-alarm-frontend:latest" \
  .
