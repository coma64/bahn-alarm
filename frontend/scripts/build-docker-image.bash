#!/usr/bin/env bash

[[ -z $1 ]] && echo 'you need to specify a version e.g. "pnpm run docker v10"' && exit 69

docker buildx build --platform linux/amd64 --push --tag "registry.zat.ong/bahn-alarm-frontend:$1" .
