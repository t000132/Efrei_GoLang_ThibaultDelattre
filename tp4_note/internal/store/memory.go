package store

import (
	"fmt"
	"sync"
	"time"

	"mini-crm/internal/models"
)

// MemoryStore implémente Storer avec un stockage en mémoire
// Utilise un mutex pour la sécurité des accès concurrents
type MemoryStore struct {
	contacts []models.Contact
	nextID   uint
	mutex    sync.RWMutex
}

// NewMemoryStore crée une nouvelle instance de MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make([]models.Contact, 0),
		nextID:   1,
	}
}

// Create ajoute un nouveau contact en mémoire
func (m *MemoryStore) Create(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Vérifier l'unicité de l'email
	for _, c := range m.contacts {
		if c.Email == contact.Email {
			return fmt.Errorf("un contact avec l'email %s existe déjà", contact.Email)
		}
	}

	contact.ID = m.nextID
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	m.contacts = append(m.contacts, *contact)
	m.nextID++

	return nil
}

// GetAll retourne tous les contacts
func (m *MemoryStore) GetAll() ([]models.Contact, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	// Copier le slice pour éviter les modifications concurrentes
	result := make([]models.Contact, len(m.contacts))
	copy(result, m.contacts)
	return result, nil
}

// GetByID récupère un contact par son ID
func (m *MemoryStore) GetByID(id uint) (*models.Contact, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, contact := range m.contacts {
		if contact.ID == id {
			// Retourner une copie pour éviter les modifications
			contactCopy := contact
			return &contactCopy, nil
		}
	}
	return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
}

// Update met à jour un contact existant
func (m *MemoryStore) Update(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, c := range m.contacts {
		if c.ID == contact.ID {
			// Vérifier l'unicité de l'email (sauf pour le contact actuel)
			for _, other := range m.contacts {
				if other.ID != contact.ID && other.Email == contact.Email {
					return fmt.Errorf("un autre contact avec l'email %s existe déjà", contact.Email)
				}
			}

			contact.UpdatedAt = time.Now()
			// Conserver la date de création originale
			contact.CreatedAt = c.CreatedAt
			m.contacts[i] = *contact
			return nil
		}
	}
	return fmt.Errorf("contact avec l'ID %d non trouvé", contact.ID)
}

// Delete supprime un contact par son ID
func (m *MemoryStore) Delete(id uint) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, contact := range m.contacts {
		if contact.ID == id {
			// Supprimer l'élément du slice
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact avec l'ID %d non trouvé", id)
}

// Close ferme le store (rien à faire pour la mémoire)
func (m *MemoryStore) Close() error {
	return nil
}
