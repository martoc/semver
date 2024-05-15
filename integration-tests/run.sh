#! /bin/bash -e

# Get the list of all .bats files in the current directory
test_files=$(ls ./integration-tests/*.bats)

# Iterate over each test file and run it
for file in $test_files; do
  echo "Running tests in $file"
  docker run -it -v "${PWD}:/workspace" \
    -e INTEGRATION_TEST_ROOT=/workspace \
    -e GOOS=linux \
    -e GOARCH=$(go env GOARCH) \
    bats/bats:latest /workspace/integration-tests/$(basename "$file")
done
