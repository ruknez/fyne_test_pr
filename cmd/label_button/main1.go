package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// название окна
	w := a.NewWindow("мое первое приложение")
	// размер окна
	w.Resize(fyne.Size{Width: 1200, Height: 600})

	// объекты с текстом которые нельзя изменить
	label := widget.NewLabel("Hello World!")
	label2 := widget.NewLabel("Label 2!")

	entry := widget.NewEntry()

	// кнопка с названием и действием
	bottom := widget.NewButton("Нажми меня", func() {
		label2.SetText(entry.Text)
	})

	c := container.NewVBox(label, label2, entry, bottom)

	w.SetContent(c)
	w.ShowAndRun()
}
