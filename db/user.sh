#!/usr/bin/env bash

run() {
  local user='go'
  local user_count="$(sudo mysql -sse "SELECT EXISTS(SELECT 1 FROM mysql.user WHERE user = '$user')")"
  local pw
  if [ "$user_count" == 0 ]; then
    pw="${1:?ERROR => must pass pw}"
    sudo mysql <<EOF
CREATE USER IF NOT EXISTS 'go'@'localhost';
ALTER USER 'go'@'localhost' IDENTIFIED BY '$pw';
EOF
  fi

  sudo mysql <<EOF
GRANT SELECT, INSERT ON one_step_gps.* TO 'go'@'localhost';
EOF
}

set -o errexit
set -o pipefail
set -o nounset
run "$@"
