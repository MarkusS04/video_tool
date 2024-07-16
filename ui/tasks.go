// Package ui stellt Funktionen bereit um eine QT Oberfläche für Video_Downloader anzuzeigen
package ui

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
)

type setupPc struct {
	Zoom struct {
		MeetingStarted        bool
		SelfUnmuteDeactivated bool
		CameraSetup           bool
		JwconfStarted         bool
	}
	Media struct {
		SongsDownloaded  bool
		SongsPlay        bool
		MediaDownloaded  bool
		RecordingStarted bool
	}
	Tablet struct {
		Started       bool
		ZoomStarted   bool
		JwconfStarted bool
	}
	Tasks struct {
		MicrofonePresent bool
		SpeakerPresent   bool
	}
}

func taskWindow(window fyne.Window) {
	var setup setupPc
	if x, err := loadTaskProgress(); err == nil {
		setup = *x
	}

	spacer := canvas.NewRectangle(color.Transparent)
	spacer.SetMinSize(fyne.NewSize(20, 20))

	zoomContainer := container.NewBorder(
		widget.NewLabel("Zoom"), // top
		nil,
		spacer,
		nil,
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Zoom Meeting gestartet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Zoom.SelfUnmuteDeactivated)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("JWConf gestartet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Zoom.JwconfStarted)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Sich selbst Lautschalten deaktiviert"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Zoom.MeetingStarted)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Kamera eingestellt"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Zoom.CameraSetup)),
			),
		),
	)

	mediaContainer := container.NewBorder(
		widget.NewLabel("Medien"), // top
		nil,
		spacer,
		nil,
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Videos heruntergeladen"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Media.MediaDownloaded)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Lieder vorbereitet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Media.SongsDownloaded)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Lieder vor/nach Versammlung abspielen\n(Zoom Freigabe)"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Media.SongsPlay)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Sonntag: Aufnahme des Vortrags"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Media.RecordingStarted)),
			),
		),
	)

	tabletContainer := container.NewBorder(
		widget.NewLabel("Tablet"), // top
		nil,
		spacer,
		nil,
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Tablet gestartet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Tablet.Started)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Zoom gestartet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Tablet.ZoomStarted)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("JWCONF gestartet"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Tablet.JwconfStarted)),
			),
		),
	)
	tasksContainer := container.NewBorder(
		widget.NewLabel("Aufgaben"), // top
		nil,
		spacer,
		nil,
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Mokrofondienst anwesend"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Tasks.MicrofonePresent)),
			),
			container.New(layout.NewGridLayoutWithColumns(2),
				widget.NewLabel("Person mit Aufgabe in Zoom anwesend\n(Vorsitzenden informieren)"),
				widget.NewCheckWithData("", binding.BindBool(&setup.Tasks.SpeakerPresent)),
			),
		),
	)

	vbox := container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			zoomContainer,
			mediaContainer,
			tabletContainer,
			tasksContainer,
		),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			backToMainMenu(window, setup.storeTaskProgress),
			widget.NewButton("Aufgaben zurücksetzen", func() {
				setup = setupPc{}
			}),
		),
	)

	window.SetContent(vbox)
	window.Show()
}

func (s *setupPc) storeTaskProgress() {
	data, err := yaml.Marshal(s)
	if err != nil {
		dialog.ShowCustom("Fehler bei YAML-Erzeugung", "OK", widget.NewLabel(err.Error()), nil)
		return
	}
	if err := os.WriteFile("taskProgress.yaml", data, 0644); err != nil {
		dialog.ShowCustom("Fehler bei YAML-Speicherung", "OK", widget.NewLabel(err.Error()), nil)
		return
	}
}

func loadTaskProgress() (*setupPc, error) {
	data, err := os.ReadFile("taskProgress.yaml")
	if err != nil {
		return nil, err
	}

	var setup setupPc
	err = yaml.Unmarshal(data, &setup)
	if err != nil {
		return nil, err
	}

	return &setup, nil
}
