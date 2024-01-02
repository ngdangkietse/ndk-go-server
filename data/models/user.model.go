package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string    `json:"fullName" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"timestamp not null default current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"timestamp not null default current_timestamp"`
}
