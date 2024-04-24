package ui

import (
	"video_tool/tools/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(app fyne.App) {
	window := app.NewWindow("Medien Downloader")
	mainMenu(window)
	window.ShowAndRun()
}

func mainMenu(window fyne.Window) {

	layout := container.NewGridWithColumns(2,
		widget.NewButton("Automatisch", func() {
			window.Hide()
			automaticDownloadMenu(window)
		}),
		widget.NewButton("Config bearbeiten", func() {
			window.Hide()
			configEditView(window)
		}),
		widget.NewButton("Musik abspielen", func() {}),
		widget.NewButton("Medien Ordner leeren", func() {
			helper.Cleanup()
			d := dialog.NewInformation("Info", "Medien Ordner erfolgreich gelehrt", window)
			d.Show()
		}),
	)

	helpButton := widget.NewButtonWithIcon("Hilfe", theme.HelpIcon(), func() {
		showInfoDialog(window)
	})

	window.SetContent(container.NewVBox(
		layout,
		helpButton,
	))

	window.Resize(fyne.NewSize(640, 400))
	window.Show()
}

func showInfoDialog(window fyne.Window) {
	infoText := "Die Lieder müssen aktuell manuell angegeben werden.\n" +
		"Der automatische Modus lädt alle Videos für die heutige Zusammenkunft und es können weitere Videos angegeben werden.\n" +
		"Im manuellen Modus müssen sämtliche Videos angegeben werden, die heruntergeladen werden sollen.\n" +
		"In der Config können bestimmte Funktionen bearbeitet werden.\n" +
		"Mittels Medien Ordner leeren, werden alle Medien entfernt die heruntergeladen wurden. Wenn autoremove aktiviert ist passiert dies bei jedem Download automatisch"

	infoLabel := widget.NewLabelWithStyle(infoText, fyne.TextAlignLeading, fyne.TextStyle{})
	infoLabel.Wrapping = fyne.TextWrapWord

	dialog.ShowCustom("Herzlich Willkommen im Medien Downloader", "OK", infoLabel, window)
}
