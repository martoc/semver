#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "New repository no tags one commit" {
  create_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "0.1.0"
}

@test "Repository with tags and multiple updates" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "1.3.0"
}
