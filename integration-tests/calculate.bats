#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "Calculate new semver" {
  run $BINARY_PATH calculate
  assert_success
  assert_not_equal $output ""
}
