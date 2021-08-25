#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

declare PROJECT

if docker-machine inspect "${PROJECT}" 1>/dev/null 2>&1; then
  exit 0
fi

# creating devbox

docker-machine create -d virtualbox \
  --virtualbox-cpu-count 2 \
  --virtualbox-memory 2048 \
  --virtualbox-share-folder "${HOME}:${HOME}" \
  --engine-insecure-registry 10.50.4.19:5000 \
  --engine-opt max-concurrent-downloads=30 \
  --engine-opt max-concurrent-uploads=30 \
  "${PROJECT}"
