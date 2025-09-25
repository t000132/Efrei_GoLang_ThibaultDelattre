package store

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"sync"
	"time"

	"mini-crm/internal/models"
)

// JSONStore implémente Storer avec persistance dans un fichier JSON
type JSONStore struct {
	filePath string
	contacts []models.Contact
	nextID   uint
	mutex    sync.RWMutex
}

// NewJSONStore crée une nouvelle instance de JSONStore
func NewJSONStore(filePath string) (*JSONStore, error) {
	store := &JSONStore{
		filePath: filePath,
		contacts: make([]models.Contact, 0),
		nextID:   1,
	}

	// Charger les données existantes
	if err := store.load(); err != nil {
		return nil, fmt.Errorf("erreur lors du chargement du fichier JSON: %v", err)
	}

	return store, nil
}

// load charge les contacts depuis le fichier JSON
func (j *JSONStore) load() error {
	// Vérifier si le fichier existe
	if _, err := os.Stat(j.filePath); os.IsNotExist(err) {
		// Le fichier n'existe pas, on part avec une liste vide
		return nil
	}

	data, err := os.ReadFile(j.filePath)
	if err != nil {
		return fmt.Errorf("erreur lecture fichier: %v", err)
	}

	// Si le fichier est vide, on part avec une liste vide
	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, &j.contacts); err != nil {
		return fmt.Errorf("erreur parsing JSON: %v", err)
	}

	// Déterminer le prochain ID
	for _, contact := range j.contacts {
		if contact.ID >= j.nextID {
			j.nextID = contact.ID + 1
		}
	}

	return nil
}

// save sauvegarde les contacts dans le fichier JSON
func (j *JSONStore) save() error {
	data, err := json.MarshalIndent(j.contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur sérialisation JSON: %v", err)
	}

	if err := os.WriteFile(j.filePath, data, fs.FileMode(0644)); err != nil {
		return fmt.Errorf("erreur écriture fichier: %v", err)
	}

	return nil
}

// Create ajoute un nouveau contact et sauvegarde
func (j *JSONStore) Create(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	j.mutex.Lock()
	defer j.mutex.Unlock()

	// Vérifier l'unicité de l'email
	for _, c := range j.contacts {
		if c.Email == contact.Email {
			return fmt.Errorf("un contact avec l'email %s existe déjà", contact.Email)
		}
	}

	contact.ID = j.nextID
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	j.contacts = append(j.contacts, *contact)
	j.nextID++

	return j.save()
}

// GetAll retourne tous les contacts
func (j *JSONStore) GetAll() ([]models.Contact, error) {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	result := make([]models.Contact, len(j.contacts))
	copy(result, j.contacts)
	return result, nil
}

// GetByID récupère un contact par son ID
func (j *JSONStore) GetByID(id uint) (*models.Contact, error) {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	for _, contact := range j.contacts {
		if contact.ID == id {
			contactCopy := contact
			return &contactCopy, nil
		}
	}
	return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
}

// Update met à jour un contact existant et sauvegarde
func (j *JSONStore) Update(contact *models.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	j.mutex.Lock()
	defer j.mutex.Unlock()

	for i, c := range j.contacts {
		if c.ID == contact.ID {
			// Vérifier l'unicité de l'email
			for _, other := range j.contacts {
				if other.ID != contact.ID && other.Email == contact.Email {
					return fmt.Errorf("un autre contact avec l'email %s existe déjà", contact.Email)
				}
			}

			contact.UpdatedAt = time.Now()
			contact.CreatedAt = c.CreatedAt
			j.contacts[i] = *contact
			return j.save()
		}
	}
	return fmt.Errorf("contact avec l'ID %d non trouvé", contact.ID)
}

// Delete supprime un contact par son ID et sauvegarde
func (j *JSONStore) Delete(id uint) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	for i, contact := range j.contacts {
		if contact.ID == id {
			j.contacts = append(j.contacts[:i], j.contacts[i+1:]...)
			return j.save()
		}
	}
	return fmt.Errorf("contact avec l'ID %d non trouvé", id)
}

// Close ferme le store (sauvegarde finale)
func (j *JSONStore) Close() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	return j.save()
}