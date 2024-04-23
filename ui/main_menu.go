package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func MainMenu(app fyne.App) {

	window := app.NewWindow("Medien Downloader")
	layout := container.NewGridWithColumns(2,
		widget.NewButton("Manuell", func() {}),
		widget.NewButton("Automatisch", func() {
			window.Hide()
			AutomaticDownloadMenu(window)
		}),
		widget.NewButton("Config bearbeiten", func() {}),
		widget.NewButton("Medien Ordner leeren", func() {}),
	)

	helpButton := widget.NewButtonWithIcon("Hilfe", theme.HelpIcon(), func() {
		showInfoDialog(window)
	})

	window.SetContent(container.NewVBox(
		layout,
		helpButton,
	))

	window.Resize(fyne.NewSize(640, 400))
	window.ShowAndRun()
}

func showInfoDialog(window fyne.Window) {
	infoText := "Die Lieder müssen aktuell manuell angegeben werden.\n" +
		"Der automatische Modus lädt alle Videos für die heutige Zusammenkunft und es können weitere Videos angegeben werden.\n" +
		"Im manuellen Modus müssen sämtliche Videos angegeben werden, die heruntergeladen werden sollen.\n" +
		"In der Config können bestimmte Funktionen bearbeitet werden.\n" +
		"Mittels Medien Ordner leeren, werden alle Lieder und Videos entfernt die heruntergeladen wurden."

	infoLabel := widget.NewLabel(infoText)
	infoLabel.Wrapping = fyne.TextWrapWord

	dialog.ShowCustom("Herzlich Willkommen im Medien Downloader", "OK", infoLabel, window)
}
