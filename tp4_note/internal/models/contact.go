package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Contact représente un contact dans notre CRM
type Contact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone     string         `json:"phone"`
	Company   string         `json:"company"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// String retourne une représentation string du contact
func (c Contact) String() string {
	return fmt.Sprintf("ID: %d | %s (%s) | %s | %s", 
		c.ID, c.Name, c.Email, c.Phone, c.Company)
}

// Validate vérifie que les champs obligatoires sont présents
func (c *Contact) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("le nom est obligatoire")
	}
	if c.Email == "" {
		return fmt.Errorf("l'email est obligatoire")
	}
	return nil
}