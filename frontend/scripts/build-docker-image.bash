#!/usr/bin/env bash

docker buildx build \
  --platform linux/amd64 \
  --push \
  --tag "registry.zat.ong/bahn-alarm-frontend:$(git describe --tags)" \
  --tag "registry.zat.ong/bahn-alarm-frontend:latest" \
  .
