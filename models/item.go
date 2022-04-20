package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ItemId      string `gorm:"primaryKey" json:"item_id"`
	ItemCode    int    `gorm:"not null" json:"item_code"`
	Description string `gorm:"not null" json:"description"`
	Quantity    string `gorm:"not null" json:"quantity"`
	OrderId     uint

	Order *Order `json:"omitempty"`
}
