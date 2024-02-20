FROM alpine:3.19.1

COPY target/semver /semver

ENTRYPOINT ["/semver"]
