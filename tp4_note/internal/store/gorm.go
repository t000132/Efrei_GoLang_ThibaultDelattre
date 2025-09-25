package store

import (
	"fmt"

	"mini-crm/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GORMStore implémente Storer avec GORM et SQLite
type GORMStore struct {
	db *gorm.DB
}

// NewGORMStore crée une nouvelle instance de GORMStore
func NewGORMStore(dbPath string) (*GORMStore, error) {
	// Ouvrir la connexion SQLite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erreur connexion base de données: %v", err)
	}

	// Migration automatique pour créer/mettre à jour les tables
	if err := db.AutoMigrate(&models.Contact{}); err != nil {
		return nil, fmt.Errorf("erreur migration base de données: %v", err)
	}

	return &GORMStore{db: db}, nil
}

// Create ajoute un nouveau contact en base
func (g *GORMStore) Create(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	// GORM gère automatiquement les timestamps et les contraintes
	result := g.db.Create(contact)
	if result.Error != nil {
		// Gérer l'erreur d'unicité de l'email
		if result.Error.Error() == "UNIQUE constraint failed: contacts.email" {
			return fmt.Errorf("un contact avec l'email %s existe déjà", contact.Email)
		}
		return fmt.Errorf("erreur création contact: %v", result.Error)
	}

	return nil
}

// GetAll récupère tous les contacts
func (g *GORMStore) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	
	result := g.db.Find(&contacts)
	if result.Error != nil {
		return nil, fmt.Errorf("erreur récupération contacts: %v", result.Error)
	}

	return contacts, nil
}

// GetByID récupère un contact par son ID
func (g *GORMStore) GetByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	
	result := g.db.First(&contact, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
		}
		return nil, fmt.Errorf("erreur récupération contact: %v", result.Error)
	}

	return &contact, nil
}

// Update met à jour un contact existant
func (g *GORMStore) Update(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	// Vérifier que le contact existe
	var existingContact models.Contact
	if err := g.db.First(&existingContact, contact.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("contact avec l'ID %d non trouvé", contact.ID)
		}
		return fmt.Errorf("erreur vérification contact: %v", err)
	}

	// Effectuer la mise à jour
	result := g.db.Save(contact)
	if result.Error != nil {
		// Gérer l'erreur d'unicité de l'email
		if result.Error.Error() == "UNIQUE constraint failed: contacts.email" {
			return fmt.Errorf("un autre contact avec l'email %s existe déjà", contact.Email)
		}
		return fmt.Errorf("erreur mise à jour contact: %v", result.Error)
	}

	return nil
}

// Delete supprime un contact par son ID
func (g *GORMStore) Delete(id uint) error {
	result := g.db.Delete(&models.Contact{}, id)
	if result.Error != nil {
		return fmt.Errorf("erreur suppression contact: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}

	return nil
}

// Close ferme la connexion à la base de données
func (g *GORMStore) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return fmt.Errorf("erreur récupération connexion SQL: %v", err)
	}
	return sqlDB.Close()
}