package repository

import (
	"github.com/densus/pos_service/pos/model/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product entity.Product) entity.Product
	Update(product entity.Product) entity.Product
	Delete(productID uint32, product entity.Product)
	All() []entity.Product
	FindByID(productID uint32) entity.Product
}

type productRepository struct {
	dbProductConnection *gorm.DB
}

func NewProductRepository(dbProductConn *gorm.DB) ProductRepository {
	return &productRepository{dbProductConnection: dbProductConn}
}

func (p *productRepository) Insert(product entity.Product) entity.Product {
	p.dbProductConnection.Save(&product)
	p.dbProductConnection.Preload("Merchant").Find(&product)
	return product
}

func (p *productRepository) Update(product entity.Product) entity.Product {
	p.dbProductConnection.Save(&product)
	p.dbProductConnection.Preload("Merchant").Find(&product)
	return product
}

func (p *productRepository) Delete(productID uint32, product entity.Product) {
	p.dbProductConnection.Where("id = ?", productID).Delete(&product)
}

func (p *productRepository) All() []entity.Product {
	var products []entity.Product
	p.dbProductConnection.Preload("Merchant").Find(&products)
	return products
}

func (p *productRepository) FindByID(productID uint32) entity.Product {
	var product entity.Product
	p.dbProductConnection.Preload("Merchant").Find(&product, productID)
	return product
}
