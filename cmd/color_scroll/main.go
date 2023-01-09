package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("LOLO")
	w.Resize(fyne.Size{Width: 600, Height: 600})

	btn := container.New(
		// все объекты становятся одного размера и накладываются друг на друга
		// на верхний план ставит последний элемент
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.RGBA{70, 128, 60, 1}),
		widget.NewButton("Click me!", func() { fmt.Println("click") }),
	)

	// контейнер для скрола
	scr := container.NewVScroll(
		container.NewVBox(
			canvas.NewRectangle(color.RGBA{70, 128, 60, 255}),

			widget.NewCheckGroup([]string{"a", "b", "c", "d", "e", "f", "g"}, func(s []string) {
				fmt.Println("s = ", s)
			}),

			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			widget.NewLabel("1"),
			btn,
		),
	)

	// устанавливакем мин размеры скрола (ширина, высота)
	scr.SetMinSize(fyne.NewSize(300, 100))

	w.SetContent(scr)
	w.ShowAndRun()
}
