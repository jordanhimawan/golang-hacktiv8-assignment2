package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId      uint       `gorm:"primaryKey" json:"order_id"`
	CustomerName string     `gorm:"not null" json:"customer_name" form:"customer_name" valid:"required~Your customer name is required"`
	OrderedAt    *time.Time `gorm:"not null" json:"ordered_at,omitempty"`
	Item         []Item     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}
