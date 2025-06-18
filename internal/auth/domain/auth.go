package domain

import (
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}

type TokenClaims struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	Type   TokenType `json:"type"`
	jwt.RegisteredClaims
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *LoginDTO) Validate() error {
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
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type RegisterDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

func (dto *RegisterDTO) Validate() error {
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
		case "admin", "teacher", "student":
		default:
			validationErrors.AddError(ErrRoleInvalid)
		}
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
}

func (dto *RefreshTokenDTO) Validate() error {
	validationErrors := NewValidationErrors()

	if strings.TrimSpace(dto.RefreshToken) == "" {
		validationErrors.AddError(ErrRefreshTokenRequired)
	}

	if validationErrors.HasErrors() {
		return validationErrors
	}

	return nil
}

type AuthResponse struct {
	User  interface{} `json:"user"`
	Token TokenPair   `json:"token"`
}

type TokenRepository interface {
	StoreRefreshToken(userID uuid.UUID, token string, expiresAt time.Time) error
	GetRefreshToken(token string) (uuid.UUID, error)
	DeleteRefreshToken(token string) error
	DeleteUserRefreshTokens(userID uuid.UUID) error
} 