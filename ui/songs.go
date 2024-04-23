package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func songsMenu(data *downloadData) fyne.CanvasObject {
	songsBox := container.NewVBox()

	addButton := widget.NewButton("Lied hinzufügen", func() {
		addInputField(data.Songs, songsBox, true)
	})

	songsBox.Add(addButton)

	return songsBox
}
