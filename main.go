package main

import (
	"video_tool/tools/helper"
	"video_tool/ui"

	"fyne.io/fyne/app"
)

func main() {
	helper.GetConfig()

	a := app.New()
	ui.MainMenu(a)
}
