#!/usr/bin/env bash
set -eo pipefail

if [ -L "${0}" ]; then
    ROOT=$(realpath -e "$(dirname "$(readlink -e "${0}")")/..")
else
    ROOT=$(realpath -e "$(dirname "${0}")/..")
fi

SPEC=${SPEC:-https://dl.kleister.eu/openapi/1.0.0-alpha1.yml}
COMMAND="docker"

if hash podman 2> /dev/null; then
    COMMAND="podman"
fi

"${COMMAND}" run --rm \
    -v "${ROOT}:/generate:Z" \
    --entrypoint swagger \
    ghcr.io/toolhippie/goswagger:latest \
    generate \
    client \
    --spec "${SPEC}" \
    --target /generate \
    --name Kleister \
    --client-package kleister \
    --model-package models \
    --default-scheme https
