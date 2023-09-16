package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base represents common fields used in all models.
type Base struct {
	UID       uuid.UUID `json:"uid,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// BeforeCreate works as a hook function for record creation. It
// runs just before the record is created.
func (base *Base) BeforeCreate(db *gorm.DB) error {
	base.UID = uuid.New()
	return nil
}
