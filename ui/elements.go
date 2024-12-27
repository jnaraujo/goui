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
	Children Element
}

func (app *App) Run() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 10, 10, 16, rl.Black)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

type Text struct {
	Content string
}
