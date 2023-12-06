package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Creating Graph window
func (c *calcView) createGraphWindow(app fyne.App) {
	c.windowGraph = app.NewWindow("Graph")
	c.windowGraph.SetContent(canvas.NewImageFromImage(c.emptyImg()))
	c.windowGraph.Resize(fyne.NewSize(float32(c.config.XWindowGraph), float32(c.config.YWindowGraph)))
	c.windowGraph.SetFixedSize(true)
	c.windowGraph.SetCloseIntercept(c.graphCloseReplace)
	c.setGraphHotKey()
}

func (c *calcView) graphCloseReplace() { // Hide instead Close
	c.windowGraph.Hide()
	c.graphHide = true
	c.Window.RequestFocus()
}

func (c *calcView) setGraphHotKey() {
	c.windowGraph.Canvas().SetOnTypedKey(c.GHhotKey)
}

func (c *calcView) graphViewHide() {
	if c.graphHide {
		c.windowGraph.Show()
		c.graphHide = false
	} else {
		c.windowGraph.Hide()
		c.graphHide = true
		c.Window.RequestFocus()
	}
}
