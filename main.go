package main

import (
	"./zserver"
)

func main() {
	s := zserver.New()
	s.Start()
}
