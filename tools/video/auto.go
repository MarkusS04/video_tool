// Package video handels the apicalls to prepare and download videos
package video

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"video_tool/tools/helper"
)

type (
	ApiData struct {
		Category Category
	}
	Category struct {
		Media []struct {
			//guid
			// LanguageAgnosticNaturalKey string
			//naturalKey
			//type
			//primaryCategory
			Title string
			//description
			//firstPublished
			//duration
			//durationFormattedHHMM
			//durationFormattedMinSec
			//tags
			Files []File
		}
	}
)

type File struct {
	ProgressiveDownloadURL string
	//checksum
	//filesize
	//modifiedDatetime
	//bitRate
	//duration
	//frameHeight
	//frameWidth
	Label string
	//frameRate
	//mimetype
	//subtitled
	//subtitles
}

// GetMediaFiles return a list of all videos with their title and download URL
func GetMediaFiles() (files []File, err error) {
	data, err := loadVideos()
	if err != nil {
		return nil, err
	}
	for _, media := range data.Category.Media {
		files = append(files,
			File{
				ProgressiveDownloadURL: extractHighestResolution(media),
				Label:                  media.Title,
			})
	}

	return files, err
}

func loadVideos() (apiData ApiData, err error) {
	request, err := http.NewRequest(http.MethodGet, helper.Config.JW.All, nil)
	if err != nil {
		return apiData, err
	}

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return apiData, err
	}
	defer response.Body.Close()

	buffer := new(bytes.Buffer)
	if _, err = buffer.ReadFrom(response.Body); err != nil {
		return apiData, err
	}

	err = json.Unmarshal(buffer.Bytes(), &apiData)

	return apiData, err
}

func extractHighestResolution(
	media struct {
		Title string
		Files []File
	},
) (highestResolutionURL string) {
	highestResolution := 0
	for _, file := range media.Files {

		resolutionStr := strings.TrimSuffix(file.Label, "p")
		resolution, err := strconv.Atoi(resolutionStr)
		if err != nil {
			continue
		}
		if resolution > highestResolution {
			highestResolution = resolution
			highestResolutionURL = file.ProgressiveDownloadURL

		}
	}

	return highestResolutionURL
}
