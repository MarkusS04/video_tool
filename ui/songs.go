package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func songsMenu(data *downloadData) fyne.CanvasObject {
	songsBox := container.NewVBox()

	addButton := widget.NewButton("Lied hinzufügen", func() {
		addInputField(data.Songs, songsBox, true)
	})

	songsBox.Add(addButton)

	return songsBox
}
