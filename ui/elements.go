package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Element interface{}
type E = Element

type Fragment struct {
	Children []Element
}
type F = Fragment

type App struct {
	Title    string
	Width    int32
	Height   int32
	Children Element

	x, y int32
}

type Text struct {
	Content  string
	X        int32
	Y        int32
	Color    rl.Color
	FontSize int32
}

type Div struct {
	Children Element
	Padding  int32
}

func (app *App) Run() {
	rl.InitWindow(app.Width, app.Height, app.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 10, 10, 16, rl.Black)

		app.render(app.Children)
		app.y = 0
		app.x = 0

		rl.EndDrawing()
	}
}

func (app *App) render(element Element) {
	switch e := element.(type) {
	case Text:
		rl.DrawText(e.Content, app.x, app.y, e.FontSize, e.Color)
		app.y += e.FontSize
	case Fragment:
		for _, elm := range e.Children {
			app.render(elm)
		}
	case Div:
		app.y += e.Padding
		app.x += e.Padding
		app.render(e.Children)
		app.y += e.Padding
		app.x += e.Padding
	default:
		fmt.Println("Unknown element.", e)
	}
}
