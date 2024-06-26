// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Sku              string    `gorm:"column:sku" json:"sku"`
	Name             string    `gorm:"column:name" json:"name"`
	Stock            int64     `gorm:"column:stock" json:"stock"`
	Price            int64     `gorm:"column:price" json:"price"`
	Image            string    `gorm:"column:image" json:"image"`
	TotalFinalPrice  int64     `gorm:"column:total_final_price" json:"total_final_price"`
	TotalNormalPrice int64     `gorm:"column:total_normal_price" json:"total_normal_price"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	CategoryID       int64     `gorm:"column:category_id" json:"category_id"`
	DiscountID       int64     `gorm:"column:discount_id" json:"discount_id"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
