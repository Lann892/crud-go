package models

import (
	"time"
)

type Person struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdatedAt:milli" json:"updatedAt"`
}

func (Person) TableName() string {
	return "persons"
}
