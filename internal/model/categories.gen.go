package model

import (
	"time"
	// validation "github.com/go-ozzo/ozzo-validation"
)

const TableNameCategory = "categories"

// Category mapped from table <categories>
type Categories struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Category's table name
func (*Categories) TableName() string {
	return TableNameCategory
}

// func (a Categories) Validate() error {
// 	return validation.ValidateStruct(&a,
// 		validation.Field(&a.ID, validation.Required),
// 		// Name cannot be empty, and the length must between 5 and 50
// 		validation.Field(&a.Name, validation.Required, validation.Length(5, 50)),
// 	)
// }
