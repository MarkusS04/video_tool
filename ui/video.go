package ui

import (
	"errors"
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

func automaticVideosList(videos *downloadData, window fyne.Window) fyne.CanvasObject {
	var err error
	videos.VideosAuto, err = video.GetMediaFiles()
	if err != nil {
		err = errors.New(
			"Es ist ein Fehler aufgetreten während der Ermittlung der Videos zum automatisierten Download.\n" +
				"Entweder nocheinmal probieren oder den manuellen Download benutzen.\n" +
				err.Error())

		dialog.ShowError(err, window)
		return nil
	}

	layout := container.NewVBox()
	var refreshLayout func()
	refreshLayout = func() {
		layout.Objects = nil // Clear the current layout
		for id := range *videos.VideosAuto {
			layout.Add(addTextField(id, videos.VideosAuto, refreshLayout))
		}
		layout.Refresh() // Refresh the layout to reflect changes
	}

	refreshLayout()
	return layout
}

func addTextField(fileID int, files *[]video.File, refreshLayout func()) fyne.CanvasObject {
	column := container.NewHBox()

	videoLabel := widget.NewLabel((*files)[fileID].Label)
	column.Add(videoLabel)

	var removeButton *widget.Button
	removeButton = widget.NewButton("Entfernen", func() {
		removeFile(files, fileID)
		refreshLayout()
	})
	column.Add(removeButton)
	return column
}

func removeFile(files *[]video.File, index int) {
	*files = append((*files)[:index], (*files)[index+1:]...)
}
