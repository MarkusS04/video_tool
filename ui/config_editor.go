package ui

import (
	"video_tool/tools/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func configEditView(window fyne.Window) {

	window.Hide()

	window.SetTitle("Medien Downloader - Config")
	window.Resize(fyne.NewSize(800, 640))
	window.CenterOnScreen()

	window.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			backToMainMenu(window),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Quelle Lieder:"),
				widget.NewEntryWithData(binding.BindString(&helper.Config.FS.Source)),
			),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Speicherort Medien:"),
				widget.NewEntryWithData(binding.BindString(&helper.Config.FS.Destination)),
			),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Medienort automatisch leeren:"),
				widget.NewCheckWithData("", binding.BindBool(&helper.Config.FS.Autoremove)),
			),
			widget.NewSeparator(),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("VLC Pfad:"),
				widget.NewEntryWithData(binding.BindString(&helper.Config.Music.Vlc)),
			),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Quelle Musik:"),
				widget.NewEntryWithData(binding.BindString(&helper.Config.Music.Source)),
			),
			container.New(
				layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Musik Player aktivieren:"),
				widget.NewCheckWithData("", binding.BindBool(&helper.Config.Music.Enabled)),
			),
			widget.NewButton("Config speichern", func() {
				helper.StoreConfig()
			}),
		),
	)

	window.Show()

}
