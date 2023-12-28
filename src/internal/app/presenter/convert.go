package presenter

import (
	"fmt"
	"reflect"
	m "smartcalc/internal/app/model"
	t "smartcalc/internal/app/tools"
	"strings"
)

// const FILENAME string = "history.json"

type convert struct {
}

// create ModelsInput with graph
func (c *convert) copyGraphForModel(v reflect.Value) (input m.ModelsInput, err bool) {
	input.ModelGraphData.Mode = 2
	input.Mode = 2
	if input.ModelGraphData.EqualValue = v.FieldByName("equalValue").String(); strings.EqualFold(input.ModelGraphData.EqualValue, "") {
		return input, true
	} else if input.ModelGraphData.XFromStr = v.FieldByName("xFromStr").String(); strings.EqualFold(input.ModelGraphData.XFromStr, "") {
		input.ModelGraphData.XFromStr = "0"
	} else if input.ModelGraphData.XToStr = v.FieldByName("xToStr").String(); strings.EqualFold(input.ModelGraphData.XToStr, "") {
		input.ModelGraphData.XToStr = "0"
	} else if input.ModelGraphData.YFromStr = v.FieldByName("yFromStr").String(); strings.EqualFold(input.ModelGraphData.YFromStr, "") {
		input.ModelGraphData.YFromStr = "0"
	} else if input.ModelGraphData.YToStr = v.FieldByName("yToStr").String(); strings.EqualFold(input.ModelGraphData.YToStr, "") {
		input.ModelGraphData.YToStr = "0"
	}
	if strings.EqualFold(input.ModelGraphData.XFromStr, input.ModelGraphData.XToStr) ||
		strings.EqualFold(input.ModelGraphData.YFromStr, input.ModelGraphData.YToStr) {
		return input, true
	}
	return
}

// create ModelsInput with equal
func (c *convert) copyEqualForModel(v reflect.Value) (input m.ModelsInput, err bool) {
	input.ModelEqualData.Mode = 1
	input.Mode = 1
	t.DbgPrint(fmt.Sprint(v))
	t.DbgPrint(fmt.Sprint(v.Field(1).String()))
	if input.ModelEqualData.EqualValue = v.Field(1).String(); strings.EqualFold(input.ModelEqualData.EqualValue, "") {
		return input, true
	} else if input.ModelEqualData.XEqualStr = v.Field(2).String(); strings.EqualFold(input.ModelEqualData.XEqualStr, "") {
		input.ModelEqualData.XEqualStr = "0"
		return input, false
	}
	return
}

// create ModelsInput with equation
func (c *convert) copyEquationForModel(v reflect.Value) (input m.ModelsInput, err bool) {
	input.ModelEquationData.Mode = 0
	input.Mode = 0
	if input.ModelEquationData.EqualValue = v.Field(1).String(); strings.EqualFold(input.ModelEquationData.EqualValue, "") {
		return input, true
	}
	return input, false
}

// converted interface to ModelsInput for Model
func (c *convert) UIToModel(in interface{}) (m.ModelsInput, bool) {
	v := reflect.ValueOf(in).Elem()
	// var r UIDataEquation
	// r = v
	// fmt.Println(&r)
	t.DbgPrint(fmt.Sprint(v))
	t.DbgPrint(fmt.Sprint("field0=", v.Field(0)))
	switch v.Field(0).Int() {
	case 0:
		return c.copyEquationForModel(v)
	case 1:
		t.DbgPrint("Choice Equal")
		return c.copyEqualForModel(v)
	case 2:
		return c.copyGraphForModel(v)
	}

	return m.ModelsInput{}, true
}

// converted modelsOutput to interface for View
func (c *convert) ModelToUI(output m.ModelsOutput) interface{} {
	switch output.Mode {
	case 0:
		t.DbgPrint(fmt.Sprint("EOUT:", output.ModelEquationResult))
		return output.ModelEquationResult
	case 1:
		t.DbgPrint(fmt.Sprint("EQLOUT:", output.ModelEquationResult))
		return output.ModelEqualResult
	case 2:
		t.DbgPrint(fmt.Sprint("GROUT:", output.ModelEquationResult))
		return output.ModelGraphResult
	}

	return &m.ModelResultEquation{
		Err:       true,
		Mode:      0,
		ResultStr: "errorPr",
	}
}
