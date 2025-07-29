package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Doctor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Skill    string `json:"skill"`
	Title    string `json:"title"`
}

// Validate checks if the doctor data is valid
func (d *Doctor) Validate() error {
	if len(d.Name) < 2 {
		return errors.New("name must be at least 2 characters")
	}
	if !strings.Contains(d.Email, "@") || len(d.Email) < 3 {
		return errors.New("invalid email")
	}
	if len(d.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

// HashPassword hashes the doctor's password
func (d *Doctor) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	d.Password = string(hashed)
	return nil
}

// VerifyPassword checks if the provided password matches the stored hash
func (d *Doctor) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password))
}
