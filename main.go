package main

import "fmt"

type App struct {
	children Element
}

func (app *App) Run() {
	fmt.Println("=== App ===")
	app.children.Render()
	fmt.Println("=== End App ===")
}

type Element interface {
	Render()
}

type E = Element

type F struct {
	children []Element
}

func (f F) Render() {
	fmt.Println("=== F ===")
	for _, child := range f.children {
		child.Render()
	}
	fmt.Println("=== End F ===")
}

type Text struct {
	Content string
}

func (t Text) Render() {
	fmt.Println("=== Text ===")
	fmt.Println(t.Content)
	fmt.Println("=== End Text ===")
}

type Button struct {
	Content Element
	OnClick func()
}

func (b Button) Render() {
	fmt.Println("=== Button ===")
	b.Content.Render()
	fmt.Println("=== End Button ===")
}

func main() {
	app := &App{
		children: F{
			[]E{
				Text{
					Content: "Hello, world!",
				},
				MyButton(),
			},
		},
	}
	app.Run()
}

func MyButton() Element {
	return Button{
		Content: F{
			[]E{
				Text{
					Content: "Click here 1",
				},
				Text{
					Content: "Click here 2",
				},
			},
		},
		OnClick: func() {
			fmt.Println("Button clicked!")
		},
	}
}
