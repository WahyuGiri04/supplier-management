package baseModel

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID              int       `json:"-" gorm:"primaryKey"`
	UUID            uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedDate     time.Time `json:"-" gorm:"autoCreateTime"`
	CreatedBy       int       `json:"-" gorm:"default:1"`
	LastUpdatedDate time.Time `json:"-" gorm:"autoUpdateTime"`
	LastUpdatedBy   int       `json:"-" gorm:"default:1"`
	IsActive        bool      `json:"is_active" gorm:"default:true"`
	IsDeleted       bool      `json:"is_deleted" gorm:"default:false"`
}
