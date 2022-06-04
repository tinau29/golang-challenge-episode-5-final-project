package model

import (
	"github.com/go-openapi/strfmt"
)

// Employee model
type Employee struct {
	ID         *int             `json:"id" example:"1"`
	Name       *string          `json:"name,omitempty" example:"tino"`
	Salary     *float64         `json:"salary,omitempty" example:"1000"`
	Address    *string          `json:"address,omitempty" example:"jl. jalan no 1"`
	IsContract *bool            `json:"is_contract" example:"false"`
	IsActive   *bool            `json:"is_active" example:"true"`
	Position   *string          `json:"position" example:"Programmer"`
	CreatedAt  *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // created at automatically inserted on post
	UpdatedAt  *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // updated at automatically changed on put or add on
}
