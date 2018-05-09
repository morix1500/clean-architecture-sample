package main

import (
	"github.com/morix1500/clean-architecture-sample/infrastructure"
)

func main() {
	b := new(infrastructure.Blog)
	b.Run()
}
