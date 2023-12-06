//go:generate fyne bundle -o data.go Icon.png

package view

import (
	d "smartcalc/internal/app/domains"

	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//-----------------------------------TYPES

// structure for output equal data
type UIDataEquation struct {
	mode       int
	equalValue string // string of equation for calc
}

type UIDataEqual struct {
	mode       int
	equalValue string // string of equal for equal or graph
	xEqualStr  string // raw data
}

// structure for output graph data
type UIDataGraph struct {
	mode                               int
	equalValue                         string // string of equal for equal or graph
	xFromStr, xToStr, yFromStr, yToStr string //raw data
}

type UIResult struct {
	err       bool   // true = error
	mode      int    // calc - 0, equal - 1 or graph - 2
	resultStr string // raw data
}

// main structure of calculator
type calcView struct {
	config *d.Cfg

	equation   string // Summary string of Users input for equation
	equalValue string // Summary string of Users input for equal

	output      *widget.Label             // widget for show input string for equation
	outputEqual *widget.Label             // widget for show input string for equal
	buttons     map[string]*widget.Button // buttons of UI
	entrys      map[string]*widget.Entry
	labels      map[string]*widget.Label
	historyList *widget.List

	Window         fyne.Window // main window
	windowGraph    fyne.Window // Graph window
	windowHistory  fyne.Window // History window
	aptabs         *container.AppTabs
	history        []d.HistoryItem
	historyWBottom fyne.Container

	activeTab   int  // Active tab indicator
	historyHide bool // History hide indicator
	graphHide   bool // Graph hide indicator

	chanwrite chan string //chanal for output data
}

//---------------------------------------Main function of view

func (c *calcView) LoadUI(app fyne.App) {
	c.entrysInit()
	c.labelsInit()
	c.buttonsInit()

	c.createGraphWindow(app)
	c.createMainWindow(app)
	c.createHistoryWindow(app, &c.history)
}

//------------------------------------NEW OBJECT

// Creating calcView object
func NewCalcView(c *d.Cfg) *calcView {
	return &calcView{
		config:      c,
		buttons:     make(map[string]*widget.Button, 100),
		entrys:      make(map[string]*widget.Entry, 10),
		labels:      make(map[string]*widget.Label, 10),
		historyHide: true,
		graphHide:   true,
	}
}

// Creating empty image for graph window
func (c *calcView) emptyImg() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, int(c.config.XWindowGraph), int(c.config.YWindowGraph)))
	for y := 0; y < int(c.config.YWindowGraph); y++ {
		for x := 0; x < int(c.config.XWindowGraph); x++ {
			img.Set(x, y, color.White)
		}
	}
	return img
}

//---------------------------------------Widgets Initiations

func (c *calcView) entrysInit() {

	entry := widget.NewEntry()
	c.entrys["xFrom"] = entry
	c.entrys["xFrom"].SetPlaceHolder("xFrom")
	c.entrys["xFrom"].SetText("-300")

	entry = widget.NewEntry()
	c.entrys["xTo"] = entry
	c.entrys["xTo"].SetPlaceHolder("xTo")
	c.entrys["xTo"].SetText("300")

	entry = widget.NewEntry()
	c.entrys["yFrom"] = entry
	c.entrys["yFrom"].SetPlaceHolder("yFrom")
	c.entrys["yFrom"].SetText("-300")

	entry = widget.NewEntry()
	c.entrys["yTo"] = entry
	c.entrys["yTo"].SetPlaceHolder("yTo")
	c.entrys["yTo"].SetText("300")

	entry = widget.NewEntry()
	c.entrys["xEqual"] = entry
	c.entrys["xEqual"].SetPlaceHolder("xValue")
	c.entrys["xEqual"].SetText("0")
}

func (c *calcView) labelsInit() {
	c.output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.output.TextStyle.Monospace = true
	// c.output.Wrapping = fyne.TextWrapWord
	c.outputEqual = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.labels["xLabel"] = widget.NewLabel("X =")
	c.labels["xLabel"].Alignment = fyne.TextAlignCenter
	c.labels["yLabel"] = widget.NewLabel("Y =")
	c.labels["yLabel"].Alignment = fyne.TextAlignCenter
	c.labels["equalLabel"] = widget.NewLabel("Equal:")
}

func (c *calcView) buttonsInit() {
	c.addButton("=", c.evaluate)
	c.buttons["="].Importance = widget.HighImportance
	c.buttons["history"] = widget.NewButton("History list (H)", c.historyViewHide)
	c.buttons["graph"] = widget.NewButton("Graph window (G)", c.graphViewHide)
}
