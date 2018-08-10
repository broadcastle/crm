package db

import "github.com/jinzhu/gorm"

// Email entry in the database.
type Email struct {
	gorm.Model

	To      string `json:"to"`
	Bcc     string `json:"bcc,omitempty"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
