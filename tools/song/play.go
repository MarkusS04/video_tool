package song

import (
	"os/exec"
	"video_tool/tools/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func PlayMusicInVlc(window fyne.Window) {
	cmd := exec.Command(helper.Config.Music.Vlc, helper.Config.Music.Source)

	err := cmd.Run()
	if err != nil {
		infoText := "Es ist folgender Fehler aufgetreten:.\n" +
			err.Error() +
			"\n\nIst der Pfad zu VLC korrekt?\n" +
			"Existiert der Ordner der Lieder?\n" +
			"Bitte prüfe die Konfiguration."

		infoLabel := widget.NewLabelWithStyle(infoText, fyne.TextAlignLeading, fyne.TextStyle{})
		infoLabel.Wrapping = fyne.TextWrapWord
		dialog.ShowCustom("Fehler bei Musikwiedergabe", "OK", infoLabel, window)
	}
}
