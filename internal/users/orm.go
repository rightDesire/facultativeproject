package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/rightDesire/facultativeproject/internal/geometry"
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

type Visit struct {
	UUID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null"`
	UserUUID     uuid.UUID  `gorm:"not null"`
	RouteUUID    uuid.UUID  `gorm:"not null"`
	LocationUUID uuid.UUID  `gorm:"not null"`
	VisitTime    time.Time  `gorm:"not null"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt    *time.Time `gorm:"type:timestamptz;default:NULL;"`
}

type Review struct {
	UUID      uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null"`
	UserUUID  uuid.UUID  `gorm:"not null"`
	RouteUUID uuid.UUID  `gorm:"not null"`
	Rating    int        `gorm:"not null"`
	Comment   string     `gorm:"type:text"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;default:NULL;"`
}

type Location struct {
	UUID        uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();not null"`
	Name        string            `gorm:"size:255;not null;uniqueIndex"`
	Description string            `gorm:"type:text"`
	Coordinates geometry.Geometry `gorm:"type:geometry(POINT,4326)"`
	CreatedAt   time.Time         `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt   time.Time         `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt   *time.Time        `gorm:"type:timestamptz;default:NULL;"`
}
