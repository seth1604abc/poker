package errors

import "fmt"

type InternalServerError struct {
	Resource string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf(`Internal Server Error: %s`, e.Resource)
}

func NewInternalServerError(resource string) error {
	return &InternalServerError{
		Resource: resource,
	}
}

// =================================
type DatabaseError struct {
	Resource string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf(`Database Error: %s`, e.Resource)
}

func NewDatabaseError(resource string) error {
	return &DatabaseError{
		Resource: resource,
	}
}

// =================================

type DuplicateAccountError struct {
	Resource string
}

func (e *DuplicateAccountError) Error() string {
	return fmt.Sprintf(`Duplicate Account Error: %s`, e.Resource)
}

func NewDuplicateAccountError(resource string) error {
	return &DuplicateAccountError{
		Resource: resource,
	}
}

// =================================

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(`Not Found Error: %s`, e.Resource)
}

func NewNotFoundError(resource string) error {
	return &NotFoundError{
		Resource: resource,
	}
}

// =================================

type PasswordValidError struct {
	Resource string
}

func (e *PasswordValidError) Error() string {
	return fmt.Sprintf(`Password Valid Error: %s`, e.Resource)
}

func NewPasswordValidError(resource string) error {
	return &PasswordValidError{
		Resource: resource,
	}
}