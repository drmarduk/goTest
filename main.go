package main

import (
	"github.com/drmarduk/goTest/animal"
	"github.com/drmarduk/goTest/feeder"
)

func main() {
	f := feeder.Feeder{Food: "mouse"}
	c := new(animal.Cat)
	f.Feed(c)
}
