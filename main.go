package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jnaraujo/goui/ui"
)

func main() {
	app := &ui.App{
		Title:    "GoUI",
		Width:    700,
		Height:   700,
		Children: Comp(),
	}

	app.Run()
}

func Comp() ui.F {
	return ui.F{
		Children: []ui.Element{
			ui.Text{
				Content:  "Hello, world!",
				X:        100,
				Y:        200,
				Color:    rl.Orange,
				FontSize: 16,
			},
			ui.Text{
				Content:  "Im not here",
				X:        130,
				Y:        240,
				Color:    rl.Blue,
				FontSize: 30,
			},
			ui.F{
				Children: []ui.Element{ui.Text{
					Content:  "Hello, world!",
					X:        10,
					Y:        32,
					Color:    rl.Orange,
					FontSize: 16,
				},
					ui.Text{
						Content:  "Im not here",
						X:        330,
						Y:        140,
						Color:    rl.Blue,
						FontSize: 30,
					}},
			},
		},
	}
}
