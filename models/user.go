package models

import (
	"errors"
	"html"
	"strings"

	lib "github.com/5amCurfew/go-gin-jwt/lib"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ///////////////////////////////////
// USER
// ///////////////////////////////////
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

func (u *User) Register() (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// BeforeCreate Hook (refer to gorm docs)
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = string(hashedPassword)

	return nil
}

// ///////////////////////////////////
// USER LOGIN
// ///////////////////////////////////
func (candidate *User) Login() (string, error) {
	u := User{}
	err := db.Model(User{}).Where("username = ?", candidate.Username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = lib.VerifyPassword(candidate.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return lib.GenerateToken(u.ID, u.IsAdmin)
}

// ///////////////////////////////////
// GET
// ///////////////////////////////////
func GetUserByID(uid uint) (User, error) {
	var u User
	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("user ID not found")
	}

	u.ClearPassword()

	return u, nil
}

func (u *User) ClearPassword() {
	u.Password = "***"
}
