package main

import (
	"image/color"

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

	// канвас для картинок
	img := canvas.NewImageFromFile("/home/dima/Pictures/Screenshot_20221217_184900.png")
	lb := widget.NewLabel("1233")

	// создаем текст через канвас для кастомизации
	canvsLabel := canvas.NewText("COLLOR", color.NRGBA{R: 23, G: 234, B: 255, A: 250})
	canvsLabe2 := canvas.NewText("COLLOR 1", color.NRGBA{R: 23, G: 234, B: 255, A: 250})
	canvsLabe3 := canvas.NewText("COLLOR 3", color.NRGBA{R: 23, G: 234, B: 255, A: 250})

	gr := container.NewGridWithColumns(2, img, container.NewVBox(canvsLabel, canvsLabe2, canvsLabe3, lb))
	w.SetContent(gr)

	w.ShowAndRun()
}
