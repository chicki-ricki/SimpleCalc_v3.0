package view

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// -------------------------------RealTime Type Handler for Calk Tab

func str255less(str string) string {
	if len(str) < 256 {
		return str
	}
	return string(str[:255])
}

// Show input data (several function)
func (c *calcView) display(newtext string) {
	c.equation = str255less(newtext)
	c.output.SetText(str255less(newtext))
}

func (c *calcView) character(char rune) {
	c.ifError()
	c.display(str255less(c.equation + string(char)))
}

func (c *calcView) text(text string) {
	c.ifError()
	c.display(str255less(c.equation + string(text)))
}

func (c *calcView) digit(d int) {
	c.ifError()
	c.character(rune(d) + '0')
}

func (c *calcView) clear() {
	c.display("")
}

func (c *calcView) backspace() {
	if len(c.equation) == 0 {
		return
	} else if c.equation == "error" {
		c.clear()
		return
	}

	c.display(c.equation[:len(c.equation)-1])
}

func (c *calcView) ifError() {
	if strings.EqualFold(c.equation, "error") {
		c.display("")
	}
}

//--------------------------------------------------------------------
// -------------------------------RealTime Type Handler for Equal and Graph Tab

func (c *calcView) displayEqual(newtext string) {
	c.equalValue = str255less(newtext)
	c.outputEqual.SetText(str255less(newtext))
}

func (c *calcView) characterEqual(char rune) {
	c.displayEqual(str255less(c.equalValue + string(char)))
}

func (c *calcView) textEqual(text string) {
	c.displayEqual(str255less(c.equalValue + string(text)))
}

func (c *calcView) digitEqual(d int) {
	c.characterEqual(rune(d) + '0')
}

func (c *calcView) clearEqual() {
	c.displayEqual("")
}

func (c *calcView) backspaceEqual() {
	if len(c.equalValue) == 0 {
		return
	} else if c.equalValue == "error" {
		c.clearEqual()
		return
	}
	c.displayEqual(c.equalValue[:len(c.equalValue)-1])
}

//-------------------------------------------------------------------------

func (c *calcView) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button
	return button
}

func (c *calcView) digitButton(number int) *widget.Button {
	str := strconv.Itoa(number)
	return c.addButton(str, func() {
		c.digit(number)
	})
}

func (c *calcView) textButton(viewText string, text string) *widget.Button {
	return c.addButton(string(viewText), func() {
		c.text(text)
	})
}

func (c *calcView) charButton(char rune) *widget.Button {
	return c.addButton(string(char), func() {
		c.character(char)
	})
}

//--------------------------------------------------------------------------
//---------------------------------------------------button creators for Equal and Graph

func (c *calcView) addButtonEqual(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button
	return button
}

func (c *calcView) digitButtonEqual(number int) *widget.Button {
	str := strconv.Itoa(number)
	return c.addButtonEqual(str, func() {
		c.digitEqual(number)
	})
}

func (c *calcView) textButtonEqual(viewText string, text string) *widget.Button {
	return c.addButtonEqual(string(viewText), func() {
		c.textEqual(text)
	})
}

func (c *calcView) charButtonEqual(char rune) *widget.Button {
	return c.addButtonEqual(string(char), func() {
		c.characterEqual(char)
	})
}

//---------------------------------------------------------------------------------

// Shortcut handlers
func (c *calcView) onPasteShortcut(shortcut fyne.Shortcut) {
	content := shortcut.(*fyne.ShortcutPaste).Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err != nil {
		return
	}
	c.display(c.equation + content)
}

func (c *calcView) onCopyShortcut(shortcut fyne.Shortcut) {
	shortcut.(*fyne.ShortcutCopy).Clipboard.SetContent(c.equation)
}

//--------------------------------------Interface for tabs composing
