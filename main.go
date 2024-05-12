package main

import (
	"fmt"
)

func main() {
	app := app{}
	err := app.initialize()
	if err != nil {
		fmt.Printf("Failed to initialize app: %s\n", err.Error())
		return
	}
	app.start()
}
