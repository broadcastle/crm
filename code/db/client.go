package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// DB is accessable database.
var DB *gorm.DB

// Contact is the main person.
type Contact struct {
	gorm.Model

	Name   string
	Email  string
	Number string

	Street string
	City   string
	State  string
	Zip    string
	Notes  []Note

	Status   Status
	StatusID uint

	NoContact bool
}

// Note has information about a contact.
type Note struct {
	gorm.Model
	ContactID uint

	Text   string
	Header string
}

// Status of the current client relationship.
type Status struct {
	gorm.Model

	ColdLead   bool
	Prospect   bool
	Evangelist bool
	Client     bool
}

// Create a contact in the database.
func (c *Contact) Create() error {
	return DB.Create(&c).Error
}

// Update a contact that is in the database.
func (c *Contact) Update(u Contact) error {
	return DB.Model(&c).Updates(&u).Error
}

// Remove a contact that is in the database.
func (c *Contact) Remove() error {
	if c.ID == 0 {
		return errors.New("need an ID")
	}

	return DB.Delete(&c).Error
}

// Query a contact given an ID.
func (c *Contact) Query() error {

	if c.ID == 0 {
		return errors.New("need an ID")
	}

	return DB.First(&c).Error

}
