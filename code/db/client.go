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

	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number"`

	Notes []Note `json:"notes"`

	Relationship   Relationship `json:"relationship"`
	RelationshipID uint         `json:"relationship_id"`

	Contacted bool `json:"contacted"`
}

// Relationship describes what this contact does.
type Relationship struct {
	gorm.Model

	Lead       bool   `json:"lead"`
	Advocate   bool   `json:"advocate"`
	Customer   bool   `json:"customer"`
	Subscriber bool   `json:"subscriber"`
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
