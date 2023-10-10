package graphic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Window() {
	a := app.New()
	w := a.NewWindow("Calc v3.0")
	width := 400
	hight := 500
	w.Resize(fyne.NewSize(float32(width), float32(hight)))

	// label := widget.NewLabel("text")
	entry := widget.NewEntry() //строка ввода
	// buttonTotal := widget.NewButton("=", func() { //кнопка=
	// 	data := entry.Text
	// 	fmt.Println(data)
	// 	label.SetText(data)
	// })
	buttonCos := widget.NewButton("cos", func() {})
	buttonAcos := widget.NewButton("acos", func() {})
	buttonCancel := widget.NewButton("C", func() {})
	buttonOpenBracket := widget.NewButton("(", func() {})
	buttonCloseBracket := widget.NewButton(")", func() {})
	buttonDiv := widget.NewButton("/", func() {})

	buttonSin := widget.NewButton("cos", func() {})
	buttonAsin := widget.NewButton("acos", func() {})
	button7 := widget.NewButton("7", func() {})
	button8 := widget.NewButton("8", func() {})
	button9 := widget.NewButton("9", func() {})
	buttonMultiply := widget.NewButton("*", func() {})

	buttonTg := widget.NewButton("cos", func() {})
	buttonAtg := widget.NewButton("acos", func() {})
	button4 := widget.NewButton("4", func() {})
	button5 := widget.NewButton("5", func() {})
	button6 := widget.NewButton("6", func() {})
	buttonSubtraction := widget.NewButton("-", func() {})

	contentColumns1 := container.NewGridWithColumns(6, buttonCos, buttonAcos, buttonCancel, buttonOpenBracket, buttonCloseBracket, buttonDiv)
	contentColumns2 := container.NewGridWithColumns(6, buttonSin, buttonAsin, button7, button8, button9, buttonMultiply)
	contentColumns3 := container.NewGridWithColumns(6, buttonTg, buttonAtg, button4, button5, button6, buttonSubtraction)
	contentColumns4 := container.NewGridWithColumns(6)
	contentColumns5 := container.NewGridWithColumns(6)
	contentRows := container.NewGridWithRows(5, contentColumns1, contentColumns2, contentColumns3, contentColumns4, contentColumns5)

	w.SetContent(container.NewVBox(
		// label,
		entry,
		// buttonTotal,
		contentRows,
	))

	w.ShowAndRun()
}
