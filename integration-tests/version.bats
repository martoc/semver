#!/usr/bin/env ./bats/bin/bats

load '/usr/lib/bats/bats-support/load'
load '/usr/lib/bats/bats-assert/load'
load 'common.sh'

@test "Get CLI Version" {
  run $BINARY_PATH version
  assert_success
  assert_equal $TAG_VERSION $(echo $output | jq -r '.version')
}
