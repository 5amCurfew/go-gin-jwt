package models

import (
	"fmt"
	"html"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// gorm Object Database
var DB *gorm.DB

// User model
type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

// Connect
func ConnectToAuthDatabase() {
	var err error

	databaseName := os.Getenv("DATABASE_NAME")
	DB, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to connect database %s", databaseName))
	} else {
		log.Infof("%s connection successful", databaseName)
	}

	DB.AutoMigrate(&User{})
}

// Authenticate user comparing candidate password with recorded (hashed) password
func (candidate *User) Login() (bool, error) {
	u := User{}
	err := DB.Model(User{}).Where("email = ?", candidate.Email).Take(&u).Error
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(u.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, err
	}

	*candidate = u
	u.ClearPassword()

	return true, nil
}

// Create a User record
func (u *User) Register() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// User BeforeCreate Hook (refer to gorm docs)
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ClearPassword() {
	u.Password = "***"
}
