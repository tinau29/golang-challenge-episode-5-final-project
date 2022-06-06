package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

// Employee model
type Employee struct {
	ID            *int             `json:"id" example:"1"`
	Name          *string          `json:"name,omitempty" example:"tino"`
	Salary        *float64         `json:"salary,omitempty" example:"1000"`
	Address       *string          `json:"address,omitempty" example:"jl. jalan no 1"`
	IsContract    *bool            `json:"is_contract" example:"false"`
	IsActive      *bool            `json:"is_active" example:"true"`
	Position      *string          `json:"position" example:"Programmer"`
	BirthOfDate   *strfmt.Date     `json:"birth_of_date" gorm:"type:date" format:"date" swaggertype:"string" example:"2000-02-22"`
	JoinOfDate    *strfmt.Date     `json:"join_of_date" gorm:"type:date" format:"date" swaggertype:"string" example:"2000-02-22"`
	AddressDetail JSONB            `json:"address_detail,omitempty" gorm:"type:jsonb;default:null"`                              //`gorm:"type:jsonb;default:'[]';not null"`
	CreatedAt     *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // created at automatically inserted on post
	UpdatedAt     *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // updated at automatically changed on put or add on
}

type AddressDetails struct {
	Number     string `json:"number,omitempty" example:"Blok A1"`
	Rt         int    `json:"rt,omitempty" example:"2"`
	Rw         int    `json:"rw,omitempty" example:"10"`
	PostalCode int    `json:"postal_code,omitempty"`
	City       string `json:"city,omitempty" example:"12345"`
	Province   string `json:"state,omitempty" example:"DKI Jakarta"`
	Country    string `json:"country,omitempty" example:"Indonesia"`
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
