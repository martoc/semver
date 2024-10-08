if [ -z "$TAG_VERSION" ]; then
  export TAG_VERSION="1.0.0"
fi

if [ -z "$GITHUB_SHA" ]; then
  export GITHUB_SHA=b65a1568cfb7cbe02a48b92859a06be5e18f0d23
fi

if [ -z "$INTEGRATION_TEST_ROOT" ]; then
  export INTEGRATION_TEST_ROOT=.
fi

BINARY_PATH="$INTEGRATION_TEST_ROOT/target/builds/semver-$GOOS-$GOARCH"

echo "TAG_VERSION=$TAG_VERSION"
echo "BINARY_PATH=$BINARY_PATH"
echo "GITHUB_SHA=$GITHUB_SHA"

apk add jq
apk add git

git config --global init.defaultBranch main

create_repository() {
  BASE=$PWD
  rm -rf .tmp/repository
  mkdir -p .tmp/repository
  cd .tmp/repository
  git init
  git checkout -b main
  git config user.email "integration-tests@build.com"
  git config user.name "Integration Test"
  cd $BASE
}

update_repository() {
  BASE=$PWD
  CHANGE_TYPE=$1
  PREFIX=$2
  if [ "$1" = "" ]; then
    CHANGE_TYPE="feat"
  fi
  cd .tmp/repository
  date >> file.txt
  git add file.txt
  git commit -m "${PREFIX}${CHANGE_TYPE}: Update file.txt"
  cd $BASE
}

tag_repository() {
  BASE=$PWD
  cd .tmp/repository
  git tag $1
  cd $BASE
}
