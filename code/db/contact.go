package db

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/nyaruka/phonenumbers"
)

// DB is accessable database.
var DB *gorm.DB

// Contact is the main person.
type Contact struct {
	gorm.Model

	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number,omitempty"`
	Slug   string `json:"slug"`

	Notes []Note `json:"notes,omitempty"`

	Lead       bool   `json:"lead,omitempty"`
	Advocate   bool   `json:"advocate,omitempty"`
	Customer   bool   `json:"customer,omitempty"`
	Subscriber bool   `json:"subscriber,omitempty"`
	Other      string `json:"other"`
	Contacted  bool   `json:"contacted"`
}

// Create a contact in the database.
func (c *Contact) Create() error {

	// Enforce the requirement of a name and email.
	if c.Name == "" || c.Email == "" {
		return errors.New("missing a name or email")
	}

	// Format the phone number
	if c.Number != "" {
		if number, err := phonenumbers.Parse(c.Number, "US"); err == nil {
			c.Number = phonenumbers.Format(number, phonenumbers.NATIONAL)
		}
	}

	// All subscribers are customers, but not all customers are subscribers.
	if c.Subscriber {
		c.Customer = true
	}

	c.Slug = slug.Make(c.Name)

	return DB.Create(&c).Error
}

// Update a contact that is in the database.
func (c *Contact) Update() error {

	// All subscribers are customers, but not all customers are subscribers.
	if c.Subscriber {
		c.Customer = true
	}

	// Format the phone number
	if c.Number != "" {
		if number, err := phonenumbers.Parse(c.Number, "US"); err == nil {
			c.Number = phonenumbers.Format(number, phonenumbers.NATIONAL)
		}
	}

	// return DB.Model(&c).Updates(&c).Error
	return DB.Save(&c).Error
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

// QueryContacts returns all of the contacts in the database.
func QueryContacts() ([]Contact, error) {

	contacts := []Contact{}

	err := DB.Find(&contacts).Error

	return contacts, err

}

// Search returns a single contact that matches c.
func (c *Contact) Search() (err error) {

	return DB.Where(&c).First(&c).Error

}

// SearchMultiple returns multiple contacts that match c.
func (c *Contact) SearchMultiple() (result []Contact, err error) {

	err = DB.Where(&c).Find(&result).Error

	return
}
