package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"primarykey; unique; type:uuid; column:id; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	Base
	Username string `gorm:"uniqueIndex;not null;size:64;" validate:"required,min=3,max=64" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;size:256;" validate:"required,email" json:"email"`
	Password string `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
}
