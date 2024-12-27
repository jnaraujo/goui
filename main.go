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
				Color:    rl.Orange,
				FontSize: 16,
			},
			ui.Text{
				Content:  "Im not here",
				Color:    rl.Blue,
				FontSize: 30,
			},
			ui.Div{
				Children: ui.F{
					Children: []ui.Element{ui.Text{
						Content:  "Hello, world!",
						Color:    rl.Orange,
						FontSize: 16,
					},
						ui.Text{
							Content:  "Im not here",
							Color:    rl.Blue,
							FontSize: 30,
						}},
				},
				Padding: 8,
			},
		},
	}
}
