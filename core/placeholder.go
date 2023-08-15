// Package core provide some functions
package core

import "log"

//go:generate ${GOPATH}/bin/mockgen -source=placeholder.go -destination=./placeholder_mock.go -package=core
type PlaceHolder interface {
	Get()
}

type Class struct {
	PlaceHolder PlaceHolder
}

// Get result
func (c *Class) Get() {
	c.PlaceHolder.Get()
	log.Println("Hello, World!")
}
