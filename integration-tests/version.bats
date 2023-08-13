#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

@test "Get CLI Version" {
  run ./target/semver version
  assert_success
  assert_not_equal $output ""
}
