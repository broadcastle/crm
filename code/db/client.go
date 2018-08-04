package db

import (
	"errors"
	"time"

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

	Notes []Note

	Relationship   Relationship
	RelationshipID uint

	Contacted bool
}

// Note has information about a contact.
type Note struct {
	gorm.Model
	ContactID uint

	Header string
	Text   string

	Task bool
	Due  time.Time

	Call  bool
	Email bool
	Event time.Time
}

// Relationship describes what this contact does.
type Relationship struct {
	gorm.Model

	Lead       bool
	Advocate   bool
	Customer   bool
	Subscriber bool
	Other      string
}

// Create a contact in the database.
func (c *Contact) Create() error {

	// Enforce the requirement of a name and email.
	if c.Name == "" || c.Email == "" {
		return errors.New("missing a name or email")
	}

	// All subscribers are customers, but not all customers are subscribers.
	if c.Relationship.Subscriber {
		c.Relationship.Customer = true
	}

	return DB.Create(&c).Error
}

// Update a contact that is in the database.
func (c *Contact) Update(u Contact) error {

	// All subscribers are customers, but not all customers are subscribers.
	if u.Relationship.Subscriber {
		u.Relationship.Customer = true
	}

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
