package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName         string    `gorm:"type:varchar(100);not null" validate:"required" json:"first_name"`
	LastName          string    `gorm:"type:varchar(100);not null" validate:"required" json:"last_name"`
	Email             string    `gorm:"type:varchar(100);not null;unique" validate:"required" json:"email"`
	Password          string    `gorm:"type:varchar(255);not null" validate:"required" json:"password,omitempty"`
	PhoneNumber       string    `gorm:"type:varchar(20);" validate:"required" json:"phone_number"`
	Address           string    `gorm:"type:text;" json:"address"`
	BirthDate         string    `gorm:"type:date;" json:"birth_date"`
	Role              string    `gorm:"type:enum('admin','user');not null" json:"role"`
	IsVerified        bool      `gorm:"default:false" json:"is_verified"`
	VerificationToken string    `gorm:"type:varchar(500);" json:"verification_token,omitempty"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}

func (I *User) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type UserSession struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserID       uint      `gorm:"not null;index"`
	Token        string    `gorm:"type:varchar(500);not null"`
	RefreshToken string    `gorm:"type:varchar(500);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (I *LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type LoginResponse struct {
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResendEmailVerifyRequest struct {
	Email string `json:"email" validate:"required"`
}

func (I *ResendEmailVerifyRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
