package db

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Note has information about a contact.
type Note struct {
	gorm.Model
	ContactID uint

	Header string `json:"header,omitempty"`
	Text   string `json:"text"`

	Task bool      `json:"task,omitempty"`
	Due  time.Time `json:"due_date,omitempty"`

	Call  bool      `json:"call,omitempty"`
	Email bool      `json:"email,omitempty"`
	Event time.Time `json:"event,omitempty"`
}

// Add a single note.
func (n *Note) Add() error {

	if n.ContactID == 0 {
		return errors.New("missing a contact ID")
	}

	return DB.Create(&n).Error
}

// Update a single note.
func (n *Note) Update(u Note) error {
	return DB.Model(&n).Updates(&u).Error
}

// Remove a single note.
func (n *Note) Remove() error {
	return DB.Delete(&n).Error
}

// Query a note given an ID.
func (n *Note) Query() error {

	if n.ID == 0 {
		return errors.New("need an ID")
	}

	return DB.First(&n).Error

}

// QueryNotes will return all of the notes.
func QueryNotes() ([]Note, error) {

	notes := []Note{}

	err := DB.Find(&notes).Error

	return notes, err

}
