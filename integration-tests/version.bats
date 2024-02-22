#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "Get CLI Version" {
  run $BINARY_PATH version
  assert_success
  assert_equal $TAG_VERSION $output
}
