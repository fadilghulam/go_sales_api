package model

import (
	"time"

	// validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-playground/validator/v10"
)

const TableNameCashier = "cashiers"

// Cashier mapped from table <cashiers>
// type Cashiers struct {
// 	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
// 	Name      string    `gorm:"column:name" json:"name"`
// 	Passcode  string    `gorm:"column:passcode" json:"passcode"`
// 	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
// 	DtmUpd    time.Time `gorm:"column:dtm_upd" json:"dtm_upd"`
// }

type Cashiers struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name" validate:"required,max=50"`
	Passcode  string    `gorm:"column:passcode" json:"passcode" validate:"required,min=8"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DtmUpd    time.Time `gorm:"column:dtm_upd" json:"dtm_upd"`
}

// TableName Cashier's table name
func (*Cashiers) TableName() string {
	return TableNameCashier
}

func ValidTimeValidator(fl validator.FieldLevel) bool {
	timeStr := fl.Field().Interface().(time.Time).Format(time.RFC3339)
	_, err := time.Parse(time.RFC3339, timeStr)
	return err == nil
}

func (cashier Cashiers) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("validTime", ValidTimeValidator)

	if cashier.CreatedAt.IsZero() {
		cashier.CreatedAt = time.Now()
	}
	if cashier.DtmUpd.IsZero() {
		cashier.DtmUpd = time.Now()
	}

	err := validate.Struct(cashier)
	if err != nil {
		return err
	}

	return nil
}

// func (a Cashiers) Validate() error {
// 	return validation.ValidateStruct(&a,
// 		// ID cannot be empty, and the length must between 5 and 50
// 		// validation.Field(&a.ID, validation.Required, ),
// 		validation.Field(&a.ID, validation.Required),
// 		// Name cannot be empty, and the length must between 5 and 50
// 		validation.Field(&a.Name, validation.Required, validation.Length(5, 50)),
// 		// Passcode cannot be empty, and must be a string consisting of two letters in upper case
// 		validation.Field(&a.Passcode, validation.Required, validation.Length(5, 50)),
// 		// DtmUpd cannot be empty, and must be a string consisting of five digits
// 		// validation.Field(&a.DtmUpd, validation.Date("2006-01-02 15:04:05")),
// 		//CreatedAt
// 		// validation.Field(&a.CreatedAt, validation.NilOrNotEmpty, validation.Date("2006-01-02 15:04:05")),
// 	)
// }
