package logger

import (
	"log"
	"os"
)

var Instance = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
