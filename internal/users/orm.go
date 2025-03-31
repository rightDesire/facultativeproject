package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null"`
	Email     string     `gorm:"size:255;not null;unique"`
	Password  string     `gorm:"size:255;not null"`
	FullName  string     `gorm:"size:255"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;default:NULL;"`
}
