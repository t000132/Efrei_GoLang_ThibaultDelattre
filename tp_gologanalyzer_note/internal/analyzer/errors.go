package analyzer

import (
	"errors"
	"fmt"
)

// Erreur fichier pas trouvé
type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

// Erreur de parsing
type ParseError struct {
	Operation string
	Err       error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur parsing %s: %v", e.Operation, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}
// Constructeurs
func NewFileNotFoundError(path string, err error) *FileNotFoundError {
	return &FileNotFoundError{
		Path: path,
		Err:  err,
	}
}

func NewParseError(operation string, err error) *ParseError {
	return &ParseError{
		Operation: operation,
		Err:       err,
	}
}

// Helpers pour vérifier le type d'erreur
func IsFileNotFound(err error) bool {
	var fileNotFoundErr *FileNotFoundError
	return errors.As(err, &fileNotFoundErr)
}

func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}