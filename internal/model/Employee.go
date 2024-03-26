package model

import (
	"database/sql"
)

type Employee struct {
	ID   int            `gorm:"column:id" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`
}
