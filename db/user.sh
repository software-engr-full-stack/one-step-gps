#!/usr/bin/env bash

run() {
  local pw="${1:?ERROR => must pass pw}"

  sudo mysql <<EOF
CREATE USER 'go'@'localhost';
GRANT SELECT, INSERT ON app.* TO 'go'@'localhost';

-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'go'@'localhost' IDENTIFIED BY '$pw';
EOF
}

set -o errexit
set -o pipefail
set -o nounset
run "$@"
