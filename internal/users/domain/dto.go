package domain

import (
	"regexp"
	"strings"
)

type CreateUserDTO struct {
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Role      UserRole `json:"role"`
}

func (dto *CreateUserDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if strings.TrimSpace(dto.Email) == "" {
		validationErrors.AddError(ErrEmailRequired)
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(dto.Email) {
			validationErrors.AddError(ErrEmailInvalid)
		}
	}

	if strings.TrimSpace(dto.Password) == "" {
		validationErrors.AddError(ErrPasswordRequired)
	} else if len(dto.Password) < 8 {
		validationErrors.AddError(ErrPasswordTooShort)
	}

	if strings.TrimSpace(dto.FirstName) == "" {
		validationErrors.AddError(ErrFirstNameRequired)
	}

	if strings.TrimSpace(dto.LastName) == "" {
		validationErrors.AddError(ErrLastNameRequired)
	}

	if dto.Role == "" {
		validationErrors.AddError(ErrRoleRequired)
	} else {
		switch dto.Role {
		case RoleAdmin, RoleTeacher, RoleStudent:
		default:
			validationErrors.AddError(ErrRoleInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type UpdateUserDTO struct {
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Role      UserRole `json:"role,omitempty"`
}

func (dto *UpdateUserDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if dto.FirstName != "" && strings.TrimSpace(dto.FirstName) == "" {
		validationErrors.AddError(ErrFirstNameRequired)
	}

	if dto.LastName != "" && strings.TrimSpace(dto.LastName) == "" {
		validationErrors.AddError(ErrLastNameRequired)
	}

	if dto.Role != "" {
		switch dto.Role {
		case RoleAdmin, RoleTeacher, RoleStudent:
		default:
			validationErrors.AddError(ErrRoleInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
} 