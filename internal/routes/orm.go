package routes

import (
	"time"

	"github.com/google/uuid"
	"github.com/rightDesire/facultativeproject/internal/geometry"
)

type Route struct {
	UUID                uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();not null"`
	RegionUUID          uuid.UUID         `gorm:"not null"`
	DifficultyLevelUUID uuid.UUID         `gorm:"not null"`
	Name                string            `gorm:"size:255;not null"`
	Description         string            `gorm:"type:text"`
	Length              float64           `gorm:"type:numeric"`
	Coordinates         geometry.Geometry `gorm:"type:geometry(LINESTRING)"`
	CreatedAt           time.Time         `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt           time.Time         `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt           *time.Time        `gorm:"type:timestamptz;default:NULL;"`
}

type RoutePoint struct {
	UUID        uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();not null"`
	RouteUUID   uuid.UUID         `gorm:"not null"`
	Name        string            `gorm:"size:255"`
	Description string            `gorm:"type:text"`
	Sequence    int               `gorm:"not null"`
	Coordinates geometry.Geometry `gorm:"type:geometry(POINT,4326)"`
}

type DifficultyLevel struct {
	UUID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null"`
	Name        string     `gorm:"size:50;not null;uniqueIndex"`
	Description string     `gorm:"type:text"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt   *time.Time `gorm:"type:timestamptz;default:NULL;"`
}

type Region struct {
	UUID        uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();not null"`
	Name        string            `gorm:"size:100;not null;uniqueIndex"`
	Description string            `gorm:"type:text"`
	Boundary    geometry.Geometry `gorm:"type:geometry(POLYGON)"`
	CreatedAt   time.Time         `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt   time.Time         `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt   *time.Time        `gorm:"type:timestamptz;default:NULL;"`
}
