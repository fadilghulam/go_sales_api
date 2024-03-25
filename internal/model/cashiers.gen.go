package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

const TableNameCashier = "cashiers"

// Cashier mapped from table <cashiers>
type Cashiers struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Passcode  string    `gorm:"column:passcode" json:"passcode"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DtmUpd    time.Time `gorm:"column:dtm_upd" json:"dtm_upd"`
}

// TableName Cashier's table name
func (*Cashiers) TableName() string {
	return TableNameCashier
}

func (a Cashiers) Validate() error {
	return validation.ValidateStruct(&a,
		// ID cannot be empty, and the length must between 5 and 50
		// validation.Field(&a.ID, validation.Required, ),
		validation.Field(&a.ID, validation.Required),
		// Name cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Name, validation.Required, validation.Length(5, 50)),
		// Passcode cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.Passcode, validation.Required, validation.Length(5, 50)),
		// DtmUpd cannot be empty, and must be a string consisting of five digits
		// validation.Field(&a.DtmUpd, validation.Date("2006-01-02 15:04:05")),
		//CreatedAt
		// validation.Field(&a.CreatedAt, validation.NilOrNotEmpty, validation.Date("2006-01-02 15:04:05")),
	)
}
