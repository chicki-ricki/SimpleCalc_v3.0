package view

import (
	d "smartcalc/internal/app/domains"
)

func (c *calcView) SetLink(link *chan string) {
	c.chanwrite = *link
}

func (c *calcView) UpdateHistory(h []d.HistoryItem) {
	c.history = h
	c.windowHistory.Canvas().Content().Refresh()
	c.historyList.ScrollToBottom()
}

func (c *calcView) GetUIData() interface{} {
	switch c.activeTab {
	case 0:
		return c.createUIDataEquation()
	case 1:
		return c.createUIDataEqual()
	case 2:
		return c.createUIDataGraph()
	}
	return "UIMode error"
}

// Display result from Presenter  ()
func (c *calcView) DisplayResult(in interface{}) {
	// c.convertToUI(in)
	result := c.convertToUIResult(in)

	if result.err {
		c.display("error")
	} else {
		switch result.mode {
		case 0, 1:
			c.display(result.resultStr)
		case 2:
			c.display(result.resultStr)
			c.showGraph()
		}
	}
	c.windowHistory.Canvas().Content().Refresh()
}
