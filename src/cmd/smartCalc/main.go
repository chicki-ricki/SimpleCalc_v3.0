// Package main launches the calculator app
//
//go:generate fyne bundle -o data.go Icon.png
package main

import (
	d "smartcalc/internal/app/domains"
	m "smartcalc/internal/app/model"
	p "smartcalc/internal/app/presenter"
	v "smartcalc/internal/app/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"fmt"
)

func main() {

	// creating application object
	app := app.New()

	// Set Theme
	if d.Config.DarkTheme != "yes" {
		app.Settings().SetTheme(theme.LightTheme())
	} else {
		app.Settings().SetTheme(theme.DarkTheme())
	}

	fmt.Println("TempFileDir:", d.Config.TempFileDir)
	fmt.Println("WorkDir:", d.Config.WorkDir)
	// set Icon
	resIconPng, err := fyne.LoadResourceFromPath(d.Config.IconPath)
	if err == nil {
		app.SetIcon(resIconPng)
	} else {
		fmt.Println(err)
	}

	// create new objects
	presenter, view, model := p.NewPresenter(), v.NewCalcView(d.Config), m.NewCalcModel(d.Config)

	//init View
	view.SetLink(&presenter.ViewDataChannel)
	view.LoadUI(app)
	view.UpdateHistory(model.GetHistory())

	// start presenters functional
	go presenter.CrossRoad(view, model)

	// show window and run application
	view.Window.ShowAndRun()
}
