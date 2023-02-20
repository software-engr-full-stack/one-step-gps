#!/usr/bin/env bash

run() {
  local table_name="${1?:ERROR => must pass table name}"
  local db_name='one_step_gps'

  sudo mysql --execute="DROP TABLE $table_name;" "$db_name"
}

set -o errexit
set -o pipefail
set -o nounset
run "$@"
