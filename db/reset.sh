#!/usr/bin/env bash

run() {
  local this_dir="$(dirname "$BASH_SOURCE")"
  sudo mysql < "$this_dir"/drop.sql &&
    sudo mysql < "$this_dir"/create.sql &&
    sudo mysql < "$this_dir/tables/customers.sql" &&
    sudo mysql < "$this_dir/data/customers.sql"
}

set -o errexit
set -o pipefail
set -o nounset
run "$@"
