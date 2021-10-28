package entity

type OutletProduct struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	OutletID uint32 `gorm:"not null" json:"outlet_id"`
	ProductID uint32 `gorm:"not null" json:"product_id"`
	Price uint32 `json:"price"`
	Product Product `gorm:"foreignKey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Outlet Outlet `gorm:"foreignKey:OutletID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}