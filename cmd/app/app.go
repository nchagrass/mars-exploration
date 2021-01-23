package main

import (
	"flag"
	"github.com/nchagrass/mars-exploration/internal/bootstrap"
)

var defaultInputPath = "./test/inputsample-1.txt"

func main() {
	var path string
	flag.StringVar(&path,
		"input-path",
		defaultInputPath,
		"default input path to read instructions from",
	)
	flag.Parse()

	bootstrap.New(path)
}
