package dto

type CreateProductDTO struct {
	Name string `form:"name" binding:"required" validate:"required,gte=5,lte=100"`
	MerchantID uint32 `form:"merchant_id" binding:"required" validate:"required"`
	//Sku string `form:"sku" binding:"required" validate:"required"`
	Image []byte
}

type UpdateProductDTO struct {
	ID uint32 `form:"id"`
	Name string `form:"name" binding:"required" validate:"required,gte=5,lte=100"`
	MerchantID uint32 `form:"merchant_id" binding:"required" validate:"required"`
	//Sku string `form:"sku" binding:"required" validate:"required,gte=5,lte=100"`
	Image []byte
}
