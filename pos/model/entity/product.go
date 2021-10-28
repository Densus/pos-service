package entity

type Product struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	MerchantID uint32 `gorm:"not null" json:"merchant_id"`
	Name string `gorm:"type:varchar(64)" json:"name"`
	Sku string `gorm:"type:varchar(16)" json:"sku"`
	Image []byte `gorm:"type:blob" json:"image"`
	Merchant Merchant `gorm:"foreignKey:MerchantID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}