package view

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func (c *calcView) createMainWindow(app fyne.App) {
	// Setting conteiners for tabs
	c.aptabs = container.NewAppTabs(
		container.NewTabItem("CALC", c.createCalcInterface()),
		container.NewTabItem("EQUAL", c.createEqualInterface()),
		container.NewTabItem("GRAPH", c.createGraphInterface()))

	// Binding open tab
	c.aptabs.OnSelected = func(ti *container.TabItem) {
		c.activeTab = c.aptabs.SelectedIndex()
	}

	// Creating programs window
	c.Window = app.NewWindow("SmartCalc")

	// Set master window
	c.Window.SetMaster()

	// Set main menu
	c.createAndSetMainMenu(app)

	// Setting content for window
	c.Window.SetContent(c.aptabs)

	// Set Window keys
	c.setMainWindowHotKey()
}

func (c *calcView) setMainWindowHotKey() {
	c.Window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		c.GHhotKey(ev)
		if ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter || string(ev.Name) == "=" {
			c.evaluate()
		} else if ev.Name == fyne.KeyBackspace && c.activeTab == 0 {
			c.backspace()
		} else if ev.Name == fyne.KeyBackspace && (c.activeTab == 1 || c.activeTab == 2) {
			c.backspaceEqual()
		} else if strings.Contains("0123456789.", string(ev.Name)) && c.activeTab == 0 {
			c.display(c.equation + string(ev.Name))
		} else if strings.Contains("0123456789.", string(ev.Name)) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + string(ev.Name))
		} else if strings.Contains("+-*/", string(ev.Name)) && c.activeTab == 0 {
			c.display(c.equation + string(ev.Name))
		} else if strings.Contains("+-*/", string(ev.Name)) {
			c.displayEqual(c.equalValue + string(ev.Name))
		} else if "E" == string(ev.Name) && c.activeTab == 0 {
			c.display(c.equation + "e")
		} else if "E" == string(ev.Name) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + "e")
		} else if "D" == string(ev.Name) && c.activeTab == 0 {
			c.display(c.equation + "^")
		} else if "D" == string(ev.Name) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + "^")
		} else if "O" == string(ev.Name) && c.activeTab == 0 {
			c.display(c.equation + "(")
		} else if "O" == string(ev.Name) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + "(")
		} else if "P" == string(ev.Name) && c.activeTab == 0 {
			c.display(c.equation + ")")
		} else if "P" == string(ev.Name) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + ")")
		} else if "X" == string(ev.Name) && (c.activeTab == 1 || c.activeTab == 2) {
			c.displayEqual(c.equalValue + "x")
		} else if strings.ContainsAny("C", string(ev.Name)) {
			if c.activeTab == 0 {
				c.clear()
			} else {
				c.clearEqual()
			}
		}
	})
}

func (c *calcView) createAndSetMainMenu(app fyne.App) {
	menuItem1 := fyne.NewMenuItem("Help", func() { c.CreateHelpWindow(app) })
	menu := fyne.NewMenu("File", menuItem1)
	mainMenu := fyne.NewMainMenu(menu)
	c.Window.SetMainMenu(mainMenu)
}

func (c *calcView) GHhotKey(ev *fyne.KeyEvent) {
	switch string(ev.Name) {
	case "G":
		c.graphViewHide()
	case "H":
		c.historyViewHide()
	}
}
