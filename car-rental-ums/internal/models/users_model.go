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
	Password          string    `gorm:"type:varchar(255);not null" validate:"required" json:"password"`
	PhoneNumber       string    `gorm:"type:varchar(20);" validate:"required" json:"phone_umber"`
	Address           string    `gorm:"type:text;" json:"address"`
	BirthDate         time.Time `gorm:"type:date;" json:"birth_date"`
	Role              string    `gorm:"type:enum('admin','user');not null" json:"role"`
	IsVerified        bool      `gorm:"default:false" json:"is_verified"`
	VerificationToken string    `gorm:"type:varchar(500);"`
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