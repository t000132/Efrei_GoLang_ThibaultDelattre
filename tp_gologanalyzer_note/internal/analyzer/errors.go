package analyzer

import (
	"errors"
	"fmt"
)

// erreur lorsque le fichier journal ne peut pas être trouvé ou accessible
type FileNotFoundError struct {
	Path string
	Err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found or inaccessible: %s", e.Path)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

// erreur lors des opérations d'analyse
type ParseError struct {
	Operation string
	Err       error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error during %s: %v", e.Operation, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

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

func IsFileNotFound(err error) bool {
	var fileNotFoundErr *FileNotFoundError
	return errors.As(err, &fileNotFoundErr)
}

func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}