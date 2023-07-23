#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

cat <<EOF
gitVersion $(git describe --tags --abbrev=0)
dockerRegistry gcr.io/sean-side
EOF
