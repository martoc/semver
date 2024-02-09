package main

import (
	"github.com/martoc/semver/cmd"
	"github.com/martoc/semver/logger"
)

func main() {
	cmd.Execute()
	logger.Close()
}
