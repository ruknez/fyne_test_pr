package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("first")
	w.Resize(fyne.Size{Width: 1200, Height: 600})

	// создали карту сюда можно добавить любой виджит но один
	card := widget.NewCard(
		"name of the card",
		"description",
		canvas.NewImageFromFile("/home/dima/Pictures/Screenshot_20221217_184900.png"),
	)

	entry := widget.NewMultiLineEntry()

	// подсказака что надо вводить
	entry.SetPlaceHolder("подсказка")

	// изначальный текст (можно редактироваться)
	entry.SetText("3123123")

	// как располагается текст
	entry.Wrapping = fyne.TextWrapBreak

	btn := widget.NewButton("Save text", func() {
		fmt.Println(entry.Text)
		// делаем логирование на каждое изменение
		entry.OnChanged = func(s string) {
			fmt.Println("on change = ", s)
		}
	})

	box := container.NewVBox(card, entry, btn)

	w.SetContent(box)

	w.ShowAndRun()
}
