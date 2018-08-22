package tui

import (
	"broadcastle.co/code/crm/code/db"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

// NoteForm creates or edits a note.
func (a *App) NoteForm(id uint) *tview.Form {

	note := db.Note{}

	if id != 0 {

		note.ID = id

		if err := note.Query(); err != nil {
			logrus.Fatal(err)
		}

	}

	c, err := db.QueryContacts()
	if err != nil {
		logrus.Fatal(err)
	}

	contacts := make(map[string]uint)
	dropContacts := []string{}

	for _, contact := range c {
		contacts[contact.Name] = contact.ID
		dropContacts = append(dropContacts, contact.Name)
	}

	return tview.NewForm().
		AddDropDown("Contact", dropContacts, 0, func(option string, optionIndex int) {
			note.ContactID = contacts[option]
		}).
		AddDropDown("Purpose", []string{"Note", "Call", "Task", "Email"}, 0, func(option string, optionIndex int) {

			note.Task = false
			note.Call = false
			note.Email = false

			switch optionIndex {
			case 1:
				note.Call = true
			case 2:
				note.Task = true
			case 3:
				note.Task = true
			}

		}).
		AddInputField("Description", note.Text, 0, nil, func(text string) {
			note.Text = text
		}).
		AddButton("Save", func() {
			defer a.Stop()

			if note.ID != 0 {
				if err := note.Update(); err != nil {
					logrus.Fatal(err)
				}
				return
			}

			saving := note

			if err := saving.Create(); err != nil {
				logrus.Fatal(err)
			}

		}).
		AddButton("Quit", func() {
			a.Stop()
		})

}
