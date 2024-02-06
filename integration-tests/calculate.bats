#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "Calculate new semver" {
  run $BINARY_PATH calculate
  assert_success
  assert_equal $output "${GITHUB_SHA:0:7}"
}
