package entity

type Merchant struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(64)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(64)" json:"email"`
}
