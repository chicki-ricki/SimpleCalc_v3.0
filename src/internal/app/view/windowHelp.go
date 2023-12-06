package view

import (
	t "smartcalc/internal/app/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *calcView) CreateHelpWindow(app fyne.App) {
	WindowHelp := app.NewWindow("Help calc")
	WindowHelp.Resize(fyne.NewSize(500, 500))
	data, err := t.ReadData(c.config.AssetsDir + "help.txt")
	var content *widget.Label

	if err != nil {
		content = widget.NewLabel("Cannot load Help information")
	} else {
		content = widget.NewLabel(string(data))
	}

	helpScroll := container.NewVScroll(content)
	helpScroll.Resize(fyne.NewSize(500, 500))
	WindowHelp.SetContent(helpScroll)
	c.setMainWindowHotKey()
	WindowHelp.Show()
}
