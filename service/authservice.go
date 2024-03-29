// service/authservice.go

package service

import (
	"errors"
	"strings"

	"github.com/rohitsmart/studio/database"
	"github.com/rohitsmart/studio/model"
	"github.com/rohitsmart/studio/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = util.ErrInvalidCredentials
	ErrDuplicateUsername  = errors.New("duplicate username")
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignUp(user *model.User) error {
	user.Username = strings.ToLower(user.Username)

	if err := s.checkDuplicateUsername(user.Username); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *AuthService) checkDuplicateUsername(username string) error {
	var existingUser model.User
	if err := database.DB.Where("LOWER(username) = ?", username).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return ErrDuplicateUsername
}

func (s *AuthService) Login(inputUser *model.User) error {
	var user model.User

	if err := database.DB.Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
