package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("init window")
	w.Resize(fyne.Size{Width: 1200, Height: 600})

	checks := widget.NewCheckGroup([]string{"1", "2", "3"}, func(s []string) {
		fmt.Println("func ", s) // логирование нажатия на чек
	})

	// надор ради группу
	radio := widget.NewRadioGroup([]string{"r1", "r2"}, func(s string) {})

	btn := widget.NewButton("Push", func() {
		fmt.Println(checks.Selected)
		fmt.Println(radio.Selected)
	})

	// доьавление ссылки на что либо
	u, _ := url.Parse("www.google.com")
	link := widget.NewHyperlink("Name", u)

	w.SetContent(container.NewVBox(container.NewHBox(checks, radio), btn, link))
	w.ShowAndRun()
}
