package note

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Create a new note.
func Create(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	note, contact := db.Note{}, db.Contact{}

	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		logrus.Fatal(err)
	}

	contact.ID = uint(id)

	if err := contact.Query(); err != nil {
		logrus.Fatal(err)
	}

	if err := questions(cmd, &note); err != nil {
		logrus.Fatal(err)
	}

	if err := db.DB.Model(&contact).Association("Notes").Append(&note).Error; err != nil {
		logrus.Fatal(err)
	}

}

func questions(cmd *cobra.Command, note *db.Note) error {

	fast, _ := cmd.Flags().GetBool("fast")

	var err error

	note.Task, err = utils.CobraInputBool(cmd, "task", "is this a task?", note.Task, fast)
	if err != nil {
		return err
	}

	note.Call, err = utils.CobraInputBool(cmd, "call", "was this a call?", note.Call, fast)
	if err != nil {
		return err
	}

	note.Email, err = utils.CobraInputBool(cmd, "email", "was this a email?", note.Email, fast)
	if err != nil {
		return err
	}

	note.Header, err = utils.CobraInput(cmd, "header", "note title", note.Header, fast, false)
	if err != nil {
		return err
	}

	note.Text, err = utils.Input("content", note.Text)
	if err != nil {
		return err
	}

	return nil

}

// Edit a note
func Edit(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	id, err := utils.UfS(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	orig, note := db.Note{}, db.Note{}
	orig.ID = id

	if err := orig.Query(); err != nil {
		logrus.Fatal(err)
	}

	if err := questions(cmd, &note); err != nil {
		logrus.Fatal(err)
	}

	if err := orig.Update(note); err != nil {
		logrus.Fatal(err)
	}

}

// Remove a note
func Remove(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Warn(err)
			break
		}

		note := db.Note{}
		note.ID = id

		if err := note.Remove(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("note #%s was removed", x)

	}

}

// View a note
func View(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	notes := []db.Note{}
	var result bytes.Buffer
	var err error

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Warn(err)
			break
		}

		note := db.Note{}
		note.ID = id

		if err := note.Query(); err != nil {
			logrus.Warn(err)
			break
		}

		notes = append(notes, note)

	}

	if len(args) < 1 {

		notes, err = db.QueryNotes()
		if err != nil {
			logrus.Info("no notes")
			return
		}

	}

	var output []byte

	output, err = json.Marshal(&notes)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Indent(&result, output, "", " "); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(string(result.Bytes()))

	dir, err := cmd.Flags().GetString("output")
	if err != nil {
		logrus.Warn(err)
		return
	}

	force, _ := cmd.Flags().GetBool("force")

	if dir != "" {

		// Create a file with the dirs.
		if _, err := os.Stat(dir); err == nil && !force {
			logrus.Warn("file exists, force flag required")
		}

		if err := ioutil.WriteFile(dir, result.Bytes(), 0644); err != nil {
			logrus.Fatal(err)
		}

	}

}
