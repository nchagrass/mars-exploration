package main

import (
	"flag"
	"marsrobot/internal/bootstrap"
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

	app, _ := bootstrap.New(path)

	app.ExecuteInstructions()
}
