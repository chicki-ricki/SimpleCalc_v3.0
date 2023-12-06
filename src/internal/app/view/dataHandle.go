package view

import (
	m "smartcalc/internal/app/model"
	t "smartcalc/internal/app/tools"

	"fmt"
	"image"
	"image/png"
	"os"
	"reflect"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

// Copy graph data from entrys to output structure
func (c *calcView) createUIDataGraph() *UIDataGraph {
	return &UIDataGraph{
		mode:       2,
		equalValue: c.equalValue,
		xFromStr:   c.entrys["xFrom"].Text,
		xToStr:     c.entrys["xTo"].Text,
		yFromStr:   c.entrys["yFrom"].Text,
		yToStr:     c.entrys["yTo"].Text,
	}
}

// Copy equation data from entrys to output structure
func (c *calcView) createUIDataEquation() *UIDataEquation {
	return &UIDataEquation{
		mode:       0,
		equalValue: c.equation,
	}
}

// Copy equal data from entrys to output structure
func (c *calcView) createUIDataEqual() *UIDataEqual {
	return &UIDataEqual{
		mode:       1,
		xEqualStr:  c.entrys["xEqual"].Text,
		equalValue: c.equalValue,
	}
}

func (c *calcView) convertToUI(in interface{}) (res UIResult) {
	// var person = (*Person)(data)
	tr := reflect.ValueOf(in)
	t.DbgPrint(fmt.Sprint("t=", tr, tr.Field(0).Bool()))
	if tr.Field(0).Bool() {
		res.err = true
		return
	} else {
		res.mode = int(tr.Field(1).Int())
		switch res.mode {
		case 0:
			fmt.Println(in, reflect.TypeOf((in)))
			eq := in.(m.ModelResultEquation)
			fmt.Println(eq, reflect.TypeOf((eq)))
		case 1, 2:
			res.resultStr = tr.Field(2).String()
			return
		}
		res.err = true
	}
	return
}

func (c *calcView) convertToUIResult(in interface{}) (res UIResult) {

	tr := reflect.ValueOf(in)
	t.DbgPrint(fmt.Sprint("t=", tr, tr.Field(0).Bool()))
	if tr.Field(0).Bool() {
		res.err = true
		return
	} else {
		res.mode = int(tr.Field(1).Int())
		switch res.mode {
		case 0, 1, 2:
			res.resultStr = tr.Field(2).String()
			return
		}
		res.err = true
	}
	return
}

func (c *calcView) showGraph() {
	if tempImg, err := c.loadImage(c.config.TempFileDir + c.config.TempGraph); !err {
		c.windowGraph.SetContent(canvas.NewImageFromImage(tempImg))
	} else if err {
		c.windowGraph.SetContent(canvas.NewImageFromImage(c.emptyImg()))
	}
	c.windowGraph.Show()
	c.graphHide = false
}

func (c *calcView) loadImage(fileName string) (im image.Image, errout bool) {
	fd, err := os.Open(fileName)
	t.DbgPrint(fmt.Sprint("OPEN"))
	if err != nil {
		fmt.Println("OPEN ERROR")
		errout = true
	}
	im, err = png.Decode(fd)
	t.DbgPrint(fmt.Sprint("DECODE"))
	if err != nil {
		fmt.Println("DECODE ERROR")
		errout = true
	}
	err = fd.Close()
	if err != nil {
		fmt.Println("CLOSE ERROR")
		errout = true
	}
	return
}

// Point output Data for calculating
func (c *calcView) evaluate() {

	if c.emptyEquationCheck() {
		return
	}
	c.chanwrite <- "ready"
}

func (c *calcView) emptyEquationCheck() bool {
	if (strings.EqualFold(c.equation, "") && c.activeTab == 0) ||
		(strings.EqualFold(c.equalValue, "") && c.activeTab == 1) ||
		(strings.EqualFold(c.equalValue, "") && c.activeTab == 2) {
		return true
	}
	return false
}
