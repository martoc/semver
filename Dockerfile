FROM scratch

COPY target/semver /bin/semver

ENTRYPOINT ["/bin/semver"]
