package entity

type Outlet struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(64)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(64)" json:"email"`
	MerchantID uint32 `gorm:"not null" json:"merchant_id"`
	Merchant Merchant `gorm:"foreignKey:MerchantID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
