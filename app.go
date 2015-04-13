package main

import (
	"flag"
	"haclient"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "default", "client name")
	flag.Parse()

	haclient.Go(name)
}
