// Package ui stellt Funktionen bereit um eine QT Oberfläche für Video_Downloader anzuzeigen
package ui

import (
	"sync"
	"video_tool/tools/song"
	"video_tool/tools/video"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type downloadData struct {
	Songs         *[]*widget.Entry
	VideosManuell *[]*widget.Entry
	VideosAuto    []video.File
}

// AutomaticDownloadMenu erstellt die GUI für den Automatischen Download
func AutomaticDownloadMenu(window fyne.Window) {

	download := downloadData{Songs: &[]*widget.Entry{}, VideosManuell: &[]*widget.Entry{}}

	window.SetTitle("Medien Downloader - Automatischer Download")
	window.Resize(fyne.NewSize(800, 640))
	window.CenterOnScreen()

	downloadBox := container.NewVBox()
	downloadBox.Add(widget.NewButton("Start Download", func() {
		var wg sync.WaitGroup
		wg.Add(3)

		progressBarSong := widget.NewProgressBar()
		progressBarManuell := widget.NewProgressBar()
		progressBarAuto := widget.NewProgressBar()
		downloadBox.Add(progressBarSong)
		downloadBox.Add(progressBarManuell)
		downloadBox.Add(progressBarAuto)

		func(wg *sync.WaitGroup) {
			defer wg.Done()
			var songs []string
			for _, song := range *download.Songs {
				songs = append(songs, song.Text)
			}

			song.Copy(songs, progressBarSong)
			progressBarSong.SetValue(100)
		}(&wg)

		func(wg *sync.WaitGroup) {
			defer wg.Done()
			var videosManuell []string
			for _, video := range *download.VideosManuell {
				videosManuell = append(videosManuell, (*video).Text)
			}

			progressBarManuell.SetValue(0)
			video.DownloadManuell(videosManuell, progressBarManuell)
			progressBarManuell.SetValue(100)
		}(&wg)

		func(wg *sync.WaitGroup) {
			defer wg.Done()
			progressBarAuto.SetValue(0)
			video.ExecDownload(download.VideosAuto, "A_", progressBarAuto)
			progressBarSong.SetValue(100)
		}(&wg)

		wg.Wait()

	}))

	vbox := widget.NewVBox(
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
