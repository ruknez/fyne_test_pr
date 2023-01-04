package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Registration")
	w.Resize(fyne.Size{Width: 600, Height: 600})
	/*
		btn := widget.NewButton("clic", func() {})
		// определяем размеры кнопки
		btn.Resize(fyne.Size{Width: 120, Height: 60})
		// сдвигаем кнопку
		btn.Move(fyne.NewPos(600, 300))

		entry := widget.NewEntry()
		entry.Resize(fyne.Size{Width: 120, Height: 60})
		//entry.Wrapping = fyne.TextWrapBreak
	*/

	label := widget.NewLabel("Регистрация")
	// установка жирного шрифта
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(200, 10))

	name := widget.NewMultiLineEntry()
	name.Resize(fyne.Size{Width: 300, Height: 40})
	name.Move(fyne.NewPos(100, 50))
	name.SetPlaceHolder("User name")

	lasrtName := widget.NewMultiLineEntry()
	lasrtName.Resize(fyne.Size{Width: 300, Height: 40})
	lasrtName.Move(fyne.NewPos(100, 90))
	lasrtName.SetPlaceHolder("Last Name")

	pass := widget.NewPasswordEntry()
	pass.Resize(fyne.Size{Width: 300, Height: 40})
	pass.Move(fyne.NewPos(100, 140))
	pass.SetPlaceHolder("Password")

	cont := container.NewWithoutLayout(label, name, lasrtName, pass)

	w.SetContent(cont)
	w.ShowAndRun()
}
