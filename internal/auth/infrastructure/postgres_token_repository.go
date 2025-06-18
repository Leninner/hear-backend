package infrastructure

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/leninner/hear-backend/internal/auth/domain"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	userDomain "github.com/leninner/hear-backend/internal/users/domain"
)

type PostgresTokenRepository struct {
	queries *db.Queries
}

func NewPostgresTokenRepository(queries *db.Queries) *PostgresTokenRepository {
	return &PostgresTokenRepository{
		queries: queries,
	}
}

func (r *PostgresTokenRepository) StoreRefreshToken(userID uuid.UUID, token string, expiresAt time.Time) error {
	ctx := context.Background()

	params := db.CreateRefreshTokenParams{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	_, err := r.queries.CreateRefreshToken(ctx, params)
	if err != nil {
		return userDomain.NewInternalError("failed to store refresh token", err)
	}

	return nil
}

func (r *PostgresTokenRepository) GetRefreshToken(token string) (uuid.UUID, error) {
	ctx := context.Background()

	userID, err := r.queries.GetRefreshToken(ctx, token)
	if err != nil {
		return uuid.Nil, domain.ErrTokenNotFound
	}

	return userID, nil
}

func (r *PostgresTokenRepository) DeleteRefreshToken(token string) error {
	ctx := context.Background()

	err := r.queries.DeleteRefreshToken(ctx, token)
	if err != nil {
		return userDomain.NewInternalError("failed to delete refresh token", err)
	}

	return nil
}

func (r *PostgresTokenRepository) DeleteUserRefreshTokens(userID uuid.UUID) error {
	ctx := context.Background()

	err := r.queries.DeleteUserRefreshTokens(ctx, userID)
	if err != nil {
		return userDomain.NewInternalError("failed to delete user refresh tokens", err)
	}

	return nil
} 