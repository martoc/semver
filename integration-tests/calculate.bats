#!/usr/bin/env ./bats/bin/bats

load '/usr/lib/bats/bats-support/load'
load '/usr/lib/bats/bats-assert/load'
load 'common.sh'

@test "New repository no tags one commit" {
  create_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "0" $(echo $output | jq -r .floating_version_major)
  assert_equal "0.1" $(echo $output | jq -r .floating_version_minor)
  assert_equal "0.1.0" $(echo $output | jq -r .next_version)
}

@test "New repository one tag one commit" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.0.0" $(echo $output | jq -r .next_version)
}

@test "Repository with tags and multiple tagged updates and one non tagged update" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.3" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.3.0" $(echo $output | jq -r .next_version)
}

@test "Repository with tags and multiple tags updates and multiple non tagged" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository
  update_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.3" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.3.0" $(echo $output | jq -r .next_version)
}

@test "Repository with one tag at the bottom and multiple non tagged commits" {
  create_repository
  update_repository && tag_repository "v0.1.0"
  update_repository
  update_repository
  update_repository
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "0" $(echo $output | jq -r .floating_version_major)
  assert_equal "0.2" $(echo $output | jq -r .floating_version_minor)
  assert_equal "0.2.0" $(echo $output | jq -r .next_version)
}

@test "Repository with nonsemantic tags and multiple tags updates and multiple non tagged" {
  create_repository
  update_repository && tag_repository "vahdfgahjsdhs"
  update_repository && tag_repository "vasdasds"
  update_repository && tag_repository "vdjhfjsdhfjd"
  update_repository
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "0" $(echo $output | jq -r .floating_version_major)
  assert_equal "0.1" $(echo $output | jq -r .floating_version_minor)
  assert_equal "0.1.0" $(echo $output | jq -r .next_version)
}

@test "Repository with tags and multiple tagged updates and one non tagged update updating patch" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "fix"
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.2" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.2.1" $(echo $output | jq -r .next_version)
}

@test "Repository with tags and multiple tagged updates and one non tagged breaking change" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "feat!"
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "2" $(echo $output | jq -r .floating_version_major)
  assert_equal "2.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "2.0.0" $(echo $output | jq -r .next_version)
}

@test "Repository with tags and multiple tagged updates and one non tagged and BREAKING CHANGE: prefix" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository && tag_repository "v1.1.0"
  update_repository && tag_repository "v1.2.0"
  update_repository "BREAKING CHANGE"
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "2" $(echo $output | jq -r .floating_version_major)
  assert_equal "2.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "2.0.0" $(echo $output | jq -r .next_version)
}

@test "New repository one tag one commit move floating tag to new commit" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1.0.0" $(echo $output | jq -r .next_version)
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.0" $(echo $output | jq -r .floating_version_minor)
  update_repository fix
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.0.1" $(echo $output | jq -r .next_version)
}

@test "Azure DevOps merge commit for a patch" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository "fix" "Merged PR 12345: "
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "1" $(echo $output | jq -r .floating_version_major)
  assert_equal "1.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "1.0.1" $(echo $output | jq -r .next_version)
}

@test "Azure DevOps merge commit for a breaking change" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository "BREAKING CHANGE" "Merged PR 12345: "
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "2" $(echo $output | jq -r .floating_version_major)
  assert_equal "2.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "2.0.0" $(echo $output | jq -r .next_version)
}

@test "Azure DevOps merge commit for a breaking change with factorial" {
  create_repository
  update_repository && tag_repository "v1.0.0"
  update_repository "refactor!" "Merged PR 12345: "
  run $BINARY_PATH calculate --path .tmp/repository --add-floating-tags
  assert_success
  assert_equal "2" $(echo $output | jq -r .floating_version_major)
  assert_equal "2.0" $(echo $output | jq -r .floating_version_minor)
  assert_equal "2.0.0" $(echo $output | jq -r .next_version)
}
