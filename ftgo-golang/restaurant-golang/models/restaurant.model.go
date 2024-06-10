package models

import "github.com/google/uuid"

type Restaurant struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name    string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	Address string    `gorm:"not null" json:"address,omitempty"`
	Menus   []Menu    `gorm:"foreignKey:RestaurantId"`
	Ticket  []Ticket  `gorm:"foreignKey:RestaurantId"`
}

type Menu struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name         string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	RestaurantId string    `gorm:"column:restaurant_id"`
	Restaurant   Restaurant
	Items        []MenuItem `gorm:"foreignKey:MenuId"`
}

type MenuItem struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name   string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	Price  float64   `gorm:"not null" json:"price,omitempty"`
	MenuId string    `gorm:"column:menu_id"`
	Menu   Menu
}
type Ticket struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Status  string    `gorm:"not null" json:"status,omitempty"`
	OrderId string    `gorm:"not null" json:"order_id,omitempty"`
	//
	RestaurantId string `gorm:"column:restaurant_id"`
	Restaurant   Restaurant
	//
}

type CreateRestaurantRequest struct {
	Name    string `json:"name"  binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateRestaurantRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type CreateMenuRequest struct {
	Name         string `json:"name"  binding:"required"`
	RestaurantId string `json:"restaurantId" binding:"required"`
}
type UpdateMenuRequest struct {
	Name         string `json:"name"  binding:"required"`
	RestaurantId string `json:"restaurantId" binding:"required"`
}
