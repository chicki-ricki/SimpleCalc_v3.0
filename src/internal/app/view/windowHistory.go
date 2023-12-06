package view

import (
	d "smartcalc/internal/app/domains"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (c *calcView) sendCleanHistory() {
	c.chanwrite <- "cleanhistory"
}

func (c *calcView) createHistoryWindow(app fyne.App, history *[]d.HistoryItem) {
	c.historyList = widget.NewList(
		func() int {
			return len(*history)
		},
		func() fyne.CanvasObject {
			return &widget.Label{}
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(
				(*history)[i].Mode + "\t :  " +
					(*history)[i].Entrys +
					(*history)[i].Equation + " = " +
					(*history)[i].Result)
		})

	historyScroll := container.NewVScroll(c.historyList)
	historyScroll.Resize(fyne.NewSize(500, 400))
	c.historyWBottom = *container.NewBorder(
		nil,
		widget.NewButton("Clear History", c.sendCleanHistory),
		nil,
		nil,
		historyScroll)

	c.windowHistory = app.NewWindow("History list")
	c.windowHistory.Resize(fyne.NewSize(500, 400))
	c.windowHistory.SetContent(&c.historyWBottom)
	// c.windowHistory.SetContent(historyScroll)
	c.windowHistory.SetCloseIntercept(c.historyCloseReplace)
	c.setHistoryHotKey()

	c.historyList.OnSelected = func(id widget.ListItemID) {
		if strings.EqualFold((*history)[id].Mode, "calc") {
			c.display((*history)[id].Equation)
			c.aptabs.SelectIndex(0)
			c.windowHistory.Canvas().Unfocus()
			c.Window.RequestFocus()
		} else if strings.EqualFold((*history)[id].Mode, "equal") {
			c.displayEqual((*history)[id].Equation)
			c.entrys["xEqual"].SetText((*history)[id].XEqual)
			c.aptabs.SelectIndex(1)
			c.windowHistory.Canvas().Unfocus()
			c.Window.RequestFocus()
		} else if strings.EqualFold((*history)[id].Mode, "graph") {
			c.displayEqual((*history)[id].Equation)
			c.entrys["xFrom"].SetText((*history)[id].XFrom)
			c.entrys["xTo"].SetText((*history)[id].XTo)
			c.entrys["yFrom"].SetText((*history)[id].YFrom)
			c.entrys["yTo"].SetText((*history)[id].YTo)
			c.aptabs.SelectIndex(2)
			c.windowHistory.Canvas().Unfocus()
			c.Window.RequestFocus()
		}
	}

	c.windowHistory.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		c.GHhotKey(ev)
	})
	// c.historyList.TypedRune(r rune) = func (e rune) {
	// 	if
	// }
	// func (c *calc) onTypedRune(r rune) {
	// 	if r == 'c' {
	// 		r = 'C' // The button is using a capital C.
	// 	}

	// 	if button, ok := c.buttons[string(r)]; ok {
	// 		button.OnTapped()
	// 	}
	// }
}

// 	c.historyList.TypedKey(ev *fyne.KeyEvent) {

// }

func (c *calcView) historyCloseReplace() { // Hide instead Close
	c.windowHistory.Hide()
	c.historyHide = true
	c.Window.RequestFocus()
}

func (c *calcView) setHistoryHotKey() {
	c.windowHistory.Canvas().SetOnTypedKey(c.GHhotKey)
}

func (c *calcView) historyViewHide() {
	if c.historyHide {
		c.windowHistory.Show()
		c.historyList.ScrollToBottom()
		c.historyHide = false
	} else {
		c.windowHistory.Hide()
		c.historyHide = true
		c.Window.RequestFocus()
	}
}
