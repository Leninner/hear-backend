package domain

import (
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

type RegisterDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken"`
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