#!/usr/bin/env bash

run() {
  local this_dir="$(dirname "$BASH_SOURCE")"
  sudo mysql < "$this_dir"/drop.sql &&
    sudo mysql < "$this_dir"/create.sql
}

set -o errexit
set -o pipefail
set -o nounset
run "$@"
