FROM scratch

ARG OS = linux
ARG ARCH = amd64
ARG TAG_VERSION = 0.0.0

COPY target/$OS-$ARCH/$TAG_VERSION/semver /bin/semver

ENTRYPOINT ["/bin/semver"]
