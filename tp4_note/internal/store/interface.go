package store

import "mini-crm/internal/models"

// Storer définit l'interface pour les opérations CRUD sur les contacts .
// Cette interface permet l'injection de dependances et le changement
// de backend de stockage sans modifier la logique métier
type Storer interface {
	// Create ajoute un nouveau contact
	Create(contact *models.Contact) error
	
	// GetAll récupère tous les contacts
	GetAll() ([]models.Contact, error)
	
	// GetByID récupère un contact par son ID
	GetByID(id uint) (*models.Contact, error)
	
	// Update met à jour un contact existant
	Update(contact *models.Contact) error
	
	// Delete supprime un contact par son ID
	Delete(id uint) error
	
	// Close ferme la connexion au stockage
	Close() error
}