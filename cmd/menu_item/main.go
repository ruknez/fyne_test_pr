package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Registration")
	w.Resize(fyne.Size{Width: 600, Height: 600})

	// главное меню -> дочернее меню -> кнопки

	file_item := fyne.NewMenuItem("New file", func() {
		fmt.Println("Do somting")
	})

	save_item := fyne.NewMenuItem("save_item", func() {
		fmt.Println("Do save")
	})

	// дополнительное меню появляющееся по наводу мышы
	save_item.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Child meny 1", nil),
	)

	// нет смысла fyne сам добавит такую кнопку
	exit_item := fyne.NewMenuItem("exit_item", func() {
		os.Exit(1)
	})

	menu := fyne.NewMenu("File", file_item, save_item, exit_item)

	exitIt := fyne.NewMenuItem("exit item", func() {
		fmt.Println("edit item")
	})
	editMenu := fyne.NewMenu("edit", exitIt)

	mainMenu := fyne.NewMainMenu(menu, editMenu)

	w.SetMainMenu(mainMenu)
	w.ShowAndRun()
}
