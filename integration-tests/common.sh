if [ -z "$VERSION" ]; then
  export VERSION="0.0.0"
fi

if [ -z "$GITHUB_SHA" ]; then
  export GITHUB_SHA=b65a1568cfb7cbe02a48b92859a06be5e18f0d23
fi

BINARY_PATH="./target/semver"

echo "VERSION=$VERSION"
echo "BINARY_PATH=$BINARY_PATH"
echo "GITHUB_SHA=$GITHUB_SHA"
