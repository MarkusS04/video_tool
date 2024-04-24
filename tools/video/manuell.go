package video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"video_tool/tools/helper"

	"fyne.io/fyne/v2/widget"
	"github.com/therecipe/qt/widgets"
)

// DownloadManuell downloads all files matching the videoname
func DownloadManuell(files []string, progressBar *widget.ProgressBar) {
	updateFileName(&files)

	ExecDownload(getMediaFiles(files), "M_", progressBar)
}

func updateFileName(files *[]string) {
extract:
	for i, f := range *files {
		res := []*regexp.Regexp{
			regexp.MustCompile(`".*(pub-[^"]*_VIDEO).*"`),
			regexp.MustCompile(`".*(docid-[^"]*_VIDEO).*"`),
			regexp.MustCompile(`.*[&,?]lank=(.*).*`),
		}

		// if a regex matches, return its value
		for _, re := range res {
			if match := re.FindStringSubmatch(f); len(match) >= 2 {
				(*files)[i] = match[1]
				continue extract
			}
		}

		// try to get value after last /
		if match := helper.FindLastOccurrence(f, '/'); match != "" {
			(*files)[i] = match
			continue extract
		}

		// Remove incompatible
		copy((*files)[i:], (*files)[i+1:]) // Shift left one index.
		(*files)[len(*files)-1] = ""       // Erase last element (write zero value).
		*files = (*files)[:len(*files)-1]  // Truncate slice.
	}
}

func getMediaFiles(files []string) (f []File) {
	errFunc := func(err error) {
		messageBox := widgets.NewQMessageBox(nil)
		messageBox.SetText(fmt.Sprintf(
			`
			<h3>Fehler beim Download</h3>
			<p>Es ist ein Fehler aufgetreten während des Downloads.</p>
			<p>%s</p>
			`, err.Error()))
		messageBox.SetStandardButtons(widgets.QMessageBox__Ok)
		messageBox.SetWindowTitle("Fehler")
		messageBox.Exec()
	}

	for _, file := range files {
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(helper.Config.JW.ApiUrl, file), nil)
		if err != nil {
			errFunc(err)
			continue
		}
		response, err := (&http.Client{}).Do(request)
		if err != nil {
			errFunc(err)
			continue
		}
		defer response.Body.Close()

		buffer := new(bytes.Buffer)
		if _, err = buffer.ReadFrom(response.Body); err != nil {
			errFunc(err)
			continue
		}

		var categoryJSON Category
		err = json.Unmarshal(buffer.Bytes(), &categoryJSON)
		if err != nil {
			errFunc(err)
			continue
		}

		for _, m := range categoryJSON.Media {
			f = append(f, File{
				ProgressiveDownloadURL: extractHighestResolution(m),
				Label:                  m.Title,
			})
		}
	}

	return f
}
