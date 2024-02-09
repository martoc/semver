#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "New repository no tags one commit" {
  create_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal "0.1.0" $output
}

@test "Repository with tags and multiple tagged updates and one non tagged update" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal "1.3.0" $output
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
  assert_equal "1.3.0" $output
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
  assert_equal "0.2.0" $output
}

@test "Repository with nonsemantic tags and multiple tags updates and multiple non tagged" {
  create_repository
  update_repository && tag_repository "vahdfgahjsdhs"
  update_repository && tag_repository "vasdasds"
  update_repository && tag_repository "vdjhfjsdhfjd"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal "0.1.0" $output
}
