package main

import "marsrobot/internal/bootstrap"

func main() {
	app, _ := bootstrap.New()

	app.ExecuteInstructions()
}