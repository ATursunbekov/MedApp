package service

import (
	"MedApp/internal/model"
	repository "MedApp/internal/repository"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateClient(client model.Client) (string, error) {
	client.Password = generatePassword(client.Password)
	id, err := a.repo.CreateClient(client)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id, true)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) CreateDoctor(doctor model.Doctor) (string, error) {
	doctor.Password = generatePassword(doctor.Password)
	id, err := a.repo.CreateDoctor(doctor)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id, false)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) LoginClient(input model.ClientInput) (string, error) {
	client, err := a.repo.LoginClient(input)
	if err != nil {
		return "", err
	}
	input.Password = generatePassword(input.Password)
	if client.Password != input.Password {
		logrus.Infof("Password does not match: db) %v | input) %v", client.Password, input.Password)
		return "", errors.New("password does not match")
	}
	accessToken, err := GenerateToken(client.ID.Hex(), true)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (a *AuthService) LoginDoctor(input model.DoctorInput) (string, error) {
	doctor, err := a.repo.LoginDoctor(input)
	if err != nil {
		return "", err
	}

	input.Password = generatePassword(input.Password)
	if doctor.Password != input.Password {
		return "", errors.New("password does not match")
	}

	accessToken, err := GenerateToken(doctor.ID.Hex(), false)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

//TODO: Password

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	// Convert SALT string to byte slice
	salt := []byte(os.Getenv("SALT"))

	// Append salt when finalizing hash
	return hex.EncodeToString(hash.Sum(salt))
}

//TODO: Access token

func GenerateToken(userID string, isClient bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"isClient": isClient,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (string, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, errors.New("invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", false, errors.New("invalid user id")
	}

	isClientVal, ok := claims["isClient"]
	if !ok {
		return "", false, errors.New("missing isClient field")
	}

	var isClient bool
	switch v := isClientVal.(type) {
	case bool:
		isClient = v
	case float64:
		isClient = v == 1
	default:
		return "", false, errors.New("invalid isClient type")
	}

	return userID, isClient, nil
}
