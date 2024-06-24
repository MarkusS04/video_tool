// Package ui stellt Funktionen bereit um eine QT Oberfläche für Video_Downloader anzuzeigen
package ui

import (
	"sync"
	"time"
	"video_tool/tools/helper"
	"video_tool/tools/song"
	"video_tool/tools/video"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type downloadData struct {
	Songs         *[]*widget.Entry
	VideosManuell *[]*widget.Entry
	VideosAuto    *[]video.File
}

// automaticDownloadMenu erstellt die GUI für den Automatischen Download
func automaticDownloadMenu(window fyne.Window) {

	var clearMedia *bool
	clearMedia = &helper.Config.FS.Autoremove

	download := downloadData{Songs: &[]*widget.Entry{}, VideosManuell: &[]*widget.Entry{}}

	window.SetTitle("Medien Downloader - Automatischer Download")
	window.Resize(fyne.NewSize(800, 640))
	window.CenterOnScreen()

	downloadBox := container.NewVBox()

	var downloadBtn *widget.Button
	downloadBtn = widget.NewButton("Start Download", func() {
		execDownload(downloadBtn, downloadBox, download, *clearMedia)
		time.Sleep(2 * time.Second)
		mainMenu(window)
	})
	downloadBox.Add(downloadBtn)

	vbox := container.New(
		layout.NewVBoxLayout(),
		backToMainMenu(window),
		widget.NewCheckWithData("Medienordner leeren", binding.BindBool(clearMedia)),
		container.NewAdaptiveGrid(2,
			songsMenu(&download),
			videoMenu(&download),
		),
		layout.NewSpacer(),
		automaticVideosList(&download),
		downloadBox,
	)

	window.SetContent(vbox)
	window.Show()
}

func execDownload(btn *widget.Button, downloadBox *fyne.Container, download downloadData, clearMedia bool) {
	btn.Disable()

	var wg sync.WaitGroup
	wg.Add(3)

	progressBarSong := widget.NewProgressBar()
	progressBarManuell := widget.NewProgressBar()
	progressBarAuto := widget.NewProgressBar()
	downloadBox.Add(progressBarSong)
	downloadBox.Add(progressBarManuell)
	downloadBox.Add(progressBarAuto)

	if helper.Config.FS.Autoremove && clearMedia {
		helper.Cleanup()
	}

	func(wg *sync.WaitGroup) {
		defer wg.Done()
		var songs []string
		for _, song := range *download.Songs {
			songs = append(songs, song.Text)
		}

		song.Copy(songs, progressBarSong)
	}(&wg)

	func(wg *sync.WaitGroup) {
		defer wg.Done()
		var videosManuell []string
		for _, video := range *download.VideosManuell {
			videosManuell = append(videosManuell, (*video).Text)
		}

		progressBarManuell.SetValue(0)
		video.DownloadManuell(videosManuell, progressBarManuell)
	}(&wg)

	func(wg *sync.WaitGroup) {
		defer wg.Done()
		progressBarAuto.SetValue(0)
		video.ExecDownload(download.VideosAuto, "A_", progressBarAuto)
	}(&wg)

	wg.Wait()
}
