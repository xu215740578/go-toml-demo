package main

import (
	"flag"
	"fmt"

	"github.com/xu215740578/go-toml-demo/config"
)

func main() {
	flag.Parse()
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
}
