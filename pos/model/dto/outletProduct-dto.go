package dto

type SetPriceDTO struct {
	Price uint32 `form:"price" binding:"required" validate:"required"`
	ProductID uint32 `form:"product_id" binding:"required" validate:"required"`
	OutletID uint32 `form:"product_id" binding:"required" validate:"required"`
}
