package tui

import (
	"strconv"

	"broadcastle.co/code/crm/code/db"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

func formContacts(a *App, form *tview.Form, contact db.Contact) {

	if contact.ID != 0 {

		if err := contact.Query(); err != nil {
			logrus.Fatal(err)
		}

	}

	drop := 0

	switch {
	case contact.Lead:
		drop = 0
	case contact.Advocate:
		drop = 1
	case contact.Customer:
		drop = 2
	case contact.Subscriber:
		drop = 3
	}

	form.
		AddInputField("Name", contact.Name, 0, nil, func(text string) {
			contact.Name = text
		}).
		AddInputField("Email", contact.Email, 0, nil, func(text string) {
			contact.Email = text
		}).
		AddInputField("Phone Number", contact.Number, 0, nil, func(text string) {
			contact.Number = text
		}).
		AddDropDown("Relationship", []string{"Lead", "Advocate", "Customer", "Subscriber"}, drop, func(option string, optionIndex int) {

			contact.Lead = false
			contact.Advocate = false
			contact.Customer = false
			contact.Subscriber = false

			switch optionIndex {
			case 0:
				contact.Lead = true
			case 1:
				contact.Advocate = true
			case 2:
				contact.Customer = true
			case 3:
				contact.Subscriber = true
				contact.Customer = true
			}
		}).
		AddCheckbox("Contacted", contact.Contacted, func(checked bool) {
			contact.Contacted = checked
		}).
		AddButton("Save", func() {

			defer a.Stop()

			if contact.ID != 0 {
				if err := contact.Update(); err != nil {
					logrus.Fatal(err)
				}
				return
			}

			saving := contact

			if err := saving.Create(); err != nil {
				logrus.Fatal(err)
			}

		}).
		AddButton("Quit", func() {
			a.Stop()
		})

}

func tableContacts(table *tview.Table, data []db.Contact) {

	rows := len(data)

	headers := []string{"ID", "Name", "Email", "Number", "Relationship", "Contacted"}

	for x := range headers {
		table.SetCellSimple(0, x, headers[x])
	}

	for r := 0; r < rows; r++ {

		relationship := "N/A"

		switch {
		case data[r].Lead:
			relationship = "Lead"
		case data[r].Advocate:
			relationship = "Advocate"
		case data[r].Customer:
			relationship = "Customer"
		case data[r].Subscriber:
			relationship = "Subscriber"
		}

		columns := []string{
			strconv.Itoa(int(data[r].ID)),
			data[r].Name,
			data[r].Email,
			data[r].Number,
			relationship,
			strconv.FormatBool(data[r].Contacted),
		}

		for c := range columns {
			table.SetCellSimple(r+1, c, columns[c])
		}

	}

}
