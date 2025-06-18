package application

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leninner/hear-backend/internal/auth/domain"
	userDomain "github.com/leninner/hear-backend/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	userRepository  userDomain.Repository
	tokenRepository domain.TokenRepository
	jwtSecret       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUseCase(userRepository userDomain.Repository, tokenRepository domain.TokenRepository, jwtSecret string) *UseCase {
	return &UseCase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		jwtSecret:       []byte(jwtSecret),
		accessTokenTTL:  15 * time.Minute,
		refreshTokenTTL: 7 * 24 * time.Hour,
	}
}

func (uc *UseCase) Login(dto *domain.LoginDTO) (*domain.AuthResponse, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.userRepository.GetByEmail(dto.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	tokenPair, err := uc.generateTokenPair(user)
	if err != nil {
		return nil, userDomain.NewInternalError("failed to generate tokens", err)
	}

	return &domain.AuthResponse{
		User:  user,
		Token: *tokenPair,
	}, nil
}

func (uc *UseCase) Register(dto *domain.RegisterDTO) (*domain.AuthResponse, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	existingUser, err := uc.userRepository.GetByEmail(dto.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, userDomain.NewInternalError("failed to hash password", err)
	}

	user := userDomain.NewUser(dto.Email, string(hashedPassword), dto.FirstName, dto.LastName, userDomain.UserRole(dto.Role))
	if err := uc.userRepository.Create(user); err != nil {
		return nil, userDomain.NewInternalError("failed to create user", err)
	}

	tokenPair, err := uc.generateTokenPair(user)
	if err != nil {
		return nil, userDomain.NewInternalError("failed to generate tokens", err)
	}

	return &domain.AuthResponse{
		User:  user,
		Token: *tokenPair,
	}, nil
}

func (uc *UseCase) RefreshToken(dto *domain.RefreshTokenDTO) (*domain.TokenPair, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	claims, err := uc.validateToken(dto.RefreshToken, domain.RefreshToken)
	if err != nil {
		return nil, err
	}

	userID, err := uc.tokenRepository.GetRefreshToken(dto.RefreshToken)
	if err != nil {
		return nil, domain.ErrTokenNotFound
	}

	if userID != claims.UserID {
		return nil, domain.ErrInvalidToken
	}

	user, err := uc.userRepository.GetByID(userID)
	if err != nil {
		return nil, userDomain.ErrUserNotFound
	}

	tokenPair, err := uc.generateTokenPair(user)
	if err != nil {
		return nil, userDomain.NewInternalError("failed to generate tokens", err)
	}

	if err := uc.tokenRepository.DeleteRefreshToken(dto.RefreshToken); err != nil {
		return nil, userDomain.NewInternalError("failed to delete old refresh token", err)
	}

	return tokenPair, nil
}

func (uc *UseCase) Logout(refreshToken string) error {
	return uc.tokenRepository.DeleteRefreshToken(refreshToken)
}

func (uc *UseCase) generateTokenPair(user *userDomain.User) (*domain.TokenPair, error) {
	now := time.Now()
	accessExp := now.Add(uc.accessTokenTTL)
	refreshExp := now.Add(uc.refreshTokenTTL)

	accessToken, err := uc.generateToken(user, domain.AccessToken, accessExp)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.generateToken(user, domain.RefreshToken, refreshExp)
	if err != nil {
		return nil, err
	}

	if err := uc.tokenRepository.StoreRefreshToken(user.ID, refreshToken, refreshExp); err != nil {
		return nil, userDomain.NewInternalError("failed to store refresh token", err)
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(uc.accessTokenTTL.Seconds()),
	}, nil
}

func (uc *UseCase) generateToken(user *userDomain.User, tokenType domain.TokenType, expiresAt time.Time) (string, error) {
	now := time.Now()
	claims := domain.TokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.jwtSecret)
}

func (uc *UseCase) validateToken(tokenString string, expectedType domain.TokenType) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return uc.jwtSecret, nil
	})

	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		if claims.Type != expectedType {
			return nil, domain.ErrInvalidToken
		}
		return claims, nil
	}

	return nil, domain.ErrInvalidToken
}

func (uc *UseCase) ValidateAccessToken(tokenString string) (*domain.TokenClaims, error) {
	return uc.validateToken(tokenString, domain.AccessToken)
} 