#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

@test "Calculate new semver" {
  run ./target/semver calculate
  assert_success
  assert_not_equal $output ""
}
