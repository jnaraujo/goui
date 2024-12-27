package main

import (
	"github.com/jnaraujo/goui/ui"
)

func main() {
	app := &ui.App{
		Children: ui.Text{
			Content: "Hello, world!",
		},
	}

	app.Run()
}
