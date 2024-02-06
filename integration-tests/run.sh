#! /bin/bash -e

# Get the list of all .bats files in the current directory
test_files=$(ls ./integration-tests/*.bats)

# Iterate over each test file and run it
for file in $test_files; do
  echo "Running tests in $file"
  ./integration-tests/bats/bin/bats "$file"
done
