// Package video handels the apicalls to prepare and download videos
package video

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"video_tool/tools/helper"

	"fyne.io/fyne/v2/widget"
	"github.com/therecipe/qt/widgets"
)

// ExecDownload downloads all files
func ExecDownload(f []File, prefix string, progressBar *widget.ProgressBar) {

	var wg sync.WaitGroup
	wg.Add(len(f))

	var progresMutex sync.Mutex
	done := 0

	for i, file := range f {
		go func(i int, file File, prefix string) {
			defer wg.Done()
			downloadFile(i, file, prefix)
			progresMutex.Lock()
			done++
			progressBar.SetValue(float64(done) / float64(len(f)))
			progresMutex.Unlock()
		}(i, file, prefix)
	}
}

func downloadFile(i int, f File, prefix string) {
	errFunc := func(err error) {
		messageBox := widgets.NewQMessageBox(nil)
		messageBox.SetText(fmt.Sprintf(
			`
			<h3>Fehler beim Download</h3>
			<p>Es ist ein Fehler aufgetreten während des Downloads von %s.</p>
			<p>%s</p>
			`, f.Label, err.Error()))
		messageBox.SetStandardButtons(widgets.QMessageBox__Ok)
		messageBox.SetWindowTitle("Fehler")
		messageBox.Exec()
	}
	fileName := helper.GetLinkPath(f.Label, i, prefix)

	file, err := os.Create(fileName)
	if err != nil {
		errFunc(err)
		return
	}
	defer file.Close()

	resp, err := http.Get(f.ProgressiveDownloadURL)
	if err != nil {
		errFunc(err)
		return
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		errFunc(err)
	}
}
