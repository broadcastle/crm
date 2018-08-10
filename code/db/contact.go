package db

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
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

	Relationship   Relationship `json:"relationship"`
	RelationshipID uint         `json:"-"`

	Contacted bool `json:"contacted"`
}

// Relationship describes what this contact does.
type Relationship struct {
	gorm.Model

	Lead       bool   `json:"lead,omitempty"`
	Advocate   bool   `json:"advocate,omitempty"`
	Customer   bool   `json:"customer,omitempty"`
	Subscriber bool   `json:"subscriber,omitempty"`
	Other      string `json:"other"`
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

	c.Slug = slug.Make(c.Name)

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

	// Remove relationship entry.
	rel := Relationship{}
	DB.Model(&c).Related(&rel)

	if err := DB.Delete(&c).Error; err != nil {
		return err
	}

	return DB.Delete(&rel).Error
}

// Fill the contact with all additional data.
func (c *Contact) Fill() error {

	var e Relationship

	if err := DB.Model(&c).Related(&e).Error; err != nil {
		return err
	}

	c.Relationship = e

	return nil

}

// Query a contact given an ID.
func (c *Contact) Query() error {

	if c.ID == 0 {
		return errors.New("need an ID")
	}

	if err := DB.First(&c).Error; err != nil {
		return err
	}

	return c.Fill()

}

// QueryContacts returns all of the contacts in the database.
func QueryContacts() ([]Contact, error) {

	contacts := []Contact{}

	err := DB.Find(&contacts).Error

	return contacts, err

}

// Search returns a single contact that matches c.
func (c *Contact) Search() (err error) {

	err = DB.Where(&c).First(&c).Error

	return c.Fill()

}

// SearchMultiple returns multiple contacts that match c.
func (c *Contact) SearchMultiple() (result []Contact, err error) {

	err = DB.Where(&c).Find(&result).Error

	for y := range result {
		err = result[y].Fill()
	}

	return
}
