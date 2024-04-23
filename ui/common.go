package ui

import (
	"errors"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func addInputField(data *[]*widget.Entry, parent *fyne.Container, numbersOnly bool) {
	inputField := widget.NewEntry()
	inputField.SetPlaceHolder("Titel eingeben")
	if numbersOnly {
		inputField.Validator = fyne.StringValidator(func(s string) error {
			if len(s) > 3 {
				return errors.ErrUnsupported
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			if num < 1 || num > 156 {
				return errors.New("Zahl außerhalb des Bereichs")
			}
			return nil
		})
	}

	*data = append(*data, inputField)

	var removeButton *widget.Button
	var con *fyne.Container
	removeButton = widget.NewButton("Entfernen", func() {
		var indexToRemove int
		for i, entry := range *data {
			if entry == inputField {
				indexToRemove = i
				break
			}
		}

		*data = append((*data)[:indexToRemove], (*data)[indexToRemove+1:]...)
		parent.Remove(con)
	})

	con = container.NewHBox(inputField, removeButton)
	(*data)[len(*data)-1] = inputField

	con.Refresh()
	parent.Add(con)
}
