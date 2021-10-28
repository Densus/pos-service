package repository

import (
	"github.com/densus/pos_service/pos/model/entity"
	"gorm.io/gorm"
)

type OutletProductRepository interface {
	SetPrice(ProductID uint32, OutletID uint32, product entity.OutletProduct) entity.OutletProduct
}

type outletProductRepository struct {
	dbConnection *gorm.DB
}

func NewOutletProductRepository(dbConn *gorm.DB) OutletProductRepository {
	return &outletProductRepository{dbConnection: dbConn}
}

func (o *outletProductRepository) SetPrice(ProductID uint32, OutletID uint32, product entity.OutletProduct) entity.OutletProduct {
	o.dbConnection.Where("product_id=? AND outlet_id=?", ProductID, OutletID).Save(&product)
	return product
}