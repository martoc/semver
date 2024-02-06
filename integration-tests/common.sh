if [ -z "$VERSION" ]; then
  VERSION="0.0.0"
fi

BINARY_PATH="./target/$(go env GOOS)-$(go env GOARCH)/$VERSION/semver"

echo "VERSION=$VERSION"
echo "BINARY_PATH=$BINARY_PATH"
