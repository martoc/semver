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

@test "New repository one tag one commit" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "1.0.0"
}

@test "Repository with tags and multiple tagged updates and one non tagged update" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "1.3.0"
}

@test "Repository with tags and multiple tags updates and multiple non tagged" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  update_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "1.3.0"
}

@test "Repository with one tag at the bottom and multiple non tagged commits" {
  create_repository
  update_repository && tag_repository "v0.1.0"
  update_repository
  update_repository
  update_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "0.2.0"
}

@test "Repository with nonsemantic tags and multiple tags updates and multiple non tagged" {
  create_repository
  update_repository && tag_repository "vahdfgahjsdhs"
  update_repository && tag_repository "vasdasds"
  update_repository && tag_repository "vdjhfjsdhfjd"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "0.1.0"
}

@test "Repository with tags and multiple tagged updates and one non tagged update updating patch" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "fix"
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "1.2.1"
}

@test "Repository with tags and multiple tagged updates and one non tagged breaking change" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "feat!"
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "2.0.0"
}

@test "Repository with tags and multiple tagged updates and one non tagged and BREAKING CHANGE: prefix" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "BREAKING CHANGE"
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "2.0.0"
}
