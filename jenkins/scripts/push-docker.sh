#!/bin/bash
set -ex
set -u
set -o pipefail

REPO=quay.io/euank/alexa-parrot:$(git rev-parse --short HEAD)
docker build -t "${REPO}" .
docker push "${REPO}"
