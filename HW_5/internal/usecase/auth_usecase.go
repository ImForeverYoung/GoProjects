package usecase

import (
	"HW_5/internal/model"
	"HW_5/internal/storage"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo *storage.Storage
}

func NewAuthUsecase(repo *storage.Storage) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

// Register hashes the password and creates a new user
func (u *AuthUsecase) Register(req model.RegisterRequest) (int, error) {
	// 1. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// 2. Create user model
	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	// 3. Save to database
	return u.repo.CreateUser(context.Background(), user) // Passing background context for now
}

// Login verifies credentials and generates a JWT token
func (u *AuthUsecase) Login(req model.LoginRequest) (string, error) {
	// 1. Find user by email
	user, err := u.repo.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// 2. Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// 3. Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign token with secret key (TODO: Move secret to env variable)
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
