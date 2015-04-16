package main

import (
	"flag"
	"github.com/rrati/etcd-ha/src/haclient"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "default", "client name")
	flag.Parse()

	haclient.Go(name)
}
