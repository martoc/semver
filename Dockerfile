FROM scratch

COPY target/semver /usr/local/bin/semver

ENTRYPOINT [ "/usr/local/bin/semver" ]
