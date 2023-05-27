package services

import (
	"fmt"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	ports "github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthUseCase struct {
	Auth           ports.AuthService
	UserRepository ports.UserRepository
}

func NewAuthService(u ports.UserRepository) ports.AuthService {
	return &AuthUseCase{
		UserRepository: u,
	}
}

func (a AuthUseCase) Login(input *models.LoginRequest) (*models.LoginResponse, error) {
	user, err := a.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, errors.UserNotFound
	}

	if err := models.ComparePasswords(user.Password, input.Password); err != nil {
		return nil, errors.InvalidLogin
	}

	token, err := a.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}

func (a AuthUseCase) GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a AuthUseCase) ValidateJWT(tokenString string) error {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.InvalidToken
	}

	return nil
}
