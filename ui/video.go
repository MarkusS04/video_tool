package ui

import (
	"video_tool/tools/video"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func videoMenu(data *downloadData) fyne.CanvasObject {
	layout := container.NewVBox()

	addButton := widget.NewButton("Video hinzufügen", func() {
		addInputField(data.VideosManuell, layout, false)
	})

	layout.Add(addButton)

	return layout
}

func automaticVideosList(videos *downloadData) fyne.CanvasObject {
	var err error
	videos.VideosAuto, err = video.GetMediaFiles()
	if err != nil {
		infoText := "Es ist ein Fehler aufgetreten während der Ermittlung der Videos zum automatisierten Download. Entweder nocheinmal probieren oder den manuellen Download benutzen.\n" +
			err.Error()

		label := widget.NewLabel(infoText)
		label.Wrapping = fyne.TextWrapWord
		dialog.ShowCustom("Fehler", "OK", label, nil)
		return nil
	}

	layout := container.NewVBox()
	for id := range videos.VideosAuto {
		layout.Add(addTextField(id, videos.VideosAuto))
	}

	return layout
}

func addTextField(fileID int, files []video.File) fyne.CanvasObject {
	column := container.NewHBox()

	videoLabel := widget.NewLabel(files[fileID].Label)
	column.Add(videoLabel)

	var removeButton *widget.Button
	removeButton = widget.NewButton("Entfernen", func() {
		files = append(files[fileID:], files[:fileID+1]...)
		videoLabel.Hide()
		removeButton.Hide()
	})
	column.Add(removeButton)
	return column
}
