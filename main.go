package main

import (
	"video_tool/tools/helper"
	"video_tool/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	helper.GetConfig()

	a := app.New()
	ui.MainMenu(a)
}
