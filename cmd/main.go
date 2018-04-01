package main

import (
	"github.com/rcliao/tachikoma/views"
)

func main() {
	drawer, _ := views.NewTerminalView()

	drawer.Draw(views.ViewData{
		Main:   views.ConvertClockToMain("20:55"),
		Text:   "hello world",
		Footer: "Hello footer",
	})
}
