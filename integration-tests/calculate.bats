#!/usr/bin/env ./bats/bin/bats

load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
load 'common.sh'

@test "Calculate new semver" {
  BASE=$PWD
  rm -rf .tmp/repository
  mkdir -p .tmp/repository
  cd .tmp/repository
  git init
  git checkout -b main
  git config user.email "integration-tests@build.com"
  git config user.name "Integration Test"
  date > file.txt
  git add file.txt
  git commit -m "feat: Initial commit"
  cd $BASE
  run $BINARY_PATH calculate --path .tmp/repository
  assert_success
  assert_equal $output "0.1.0"
}
