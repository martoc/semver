FROM scratch

ARG TARGETOS
ARG TARGETARCH

COPY target/builds/semver-$TARGETOS-$TARGETARCH /usr/local/bin/semver

ENTRYPOINT [ "/usr/local/bin/semver" ]
