// Package song handels songs
package song

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"video_tool/tools/helper"

	"fyne.io/fyne/v2/widget"
	"github.com/therecipe/qt/widgets"
)

// Copy copies all songs from definied source in config to definied destination
func Copy(songs []string, progressBar *widget.ProgressBar) {
	var wg sync.WaitGroup
	wg.Add(len(songs))

	var progresMutex sync.Mutex
	done := 0

	for i, song := range songs {
		go func(song string, i int) {
			defer wg.Done()
			copySong(song, i)

			progresMutex.Lock()
			done++
			progressBar.SetValue(float64(done) / float64(len(songs)))
			progressBar.Refresh()
			progresMutex.Unlock()
		}(song, i)
	}

	wg.Wait()
	progressBar.SetValue(1)
}

func copySong(song string, count int) {
	errFunc := func(err error) {
		messageBox := widgets.NewQMessageBox(nil)
		messageBox.SetText(fmt.Sprintf(
			`
			<h3>Fehler beim Kopieren</h3>
			<p>Es ist ein Fehler aufgetreten während des Kopierens aufgetreten.</p>
			<p>%s</p>
			`, err.Error()))
		messageBox.SetStandardButtons(widgets.QMessageBox__Ok)
		messageBox.SetWindowTitle("Fehler")
		messageBox.Exec()
	}

	song, err := formatNumber(song)
	if err != nil {
		errFunc(err)
		return
	}

	orig := getOrigPath(song)
	dest := helper.GetLinkPath(song, count, "")

	// open source
	inputFile, err := os.Open(orig)
	if err != nil {
		errFunc(err)
		return
	}
	defer inputFile.Close()

	// create target
	outputFile, err := os.Create(dest)
	if err != nil {
		errFunc(err)
		return
	}
	defer outputFile.Close()

	// copy source to target
	if _, err = io.Copy(outputFile, bufio.NewReader(inputFile)); err != nil {
		errFunc(err)
	}
}

func getOrigPath(name string) string {
	return filepath.FromSlash(fmt.Sprintf("%s/%s.mp4", helper.Config.FS.Source, name))
}

func formatNumber(numStr string) (string, error) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", err
	}
	if num < 1 || num > 156 {
		return "", errors.New("Zahl außerhalb des erlaubten Bereichs")
	}
	return fmt.Sprintf("%03d", num), nil
}
