package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required,min=0"`
	Quantity    int       `json:"quantity" binding:"required,min=0"`
	ImagePath   string    `json:"imagePath"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	ImagePath   string    `json:"imagePath"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	// Any pre-save validation or modification can go here
	return nil
}

func (p *Product) ToResponse() ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		ImagePath:   p.ImagePath,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
} 
