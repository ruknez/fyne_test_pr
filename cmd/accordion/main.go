package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accardio")
	w.Resize(fyne.Size{Width: 600, Height: 600})

	// Для вставки текста из Lorem ipsum -> ctrl + shift +p
	text := "Voluptate ullamco incididunt commodo et enim Lorem in aliquip quis. Adipisicing id ad commodo culpa eiusmod exercitation cupidatat qui ipsum consequat et anim. In qui quis aute ex pariatur velit excepteur. Ad laborum duis duis occaecat fugiat mollit pariatur quis et."
	text_label := widget.NewLabel(text)
	// что бы текст адаптировался по размерам
	text_label.Wrapping = fyne.TextWrapBreak
	// первый итем в акардионе
	item1 := widget.NewAccordionItem(
		"Accordion Item 1",
		text_label,
	)

	item2 := widget.NewAccordionItem("Button", widget.NewButton("sey hello", func() { fmt.Println("LOLO") }))

	//  сам  акордион
	accordion := widget.NewAccordion(item1, item2)
	w.SetContent(accordion)
	w.ShowAndRun()
	a.Run()
}
