package service

import (
	"fmt"
	"github.com/densus/pos_service/pos/model/dto"
	"github.com/densus/pos_service/pos/model/entity"
	"github.com/densus/pos_service/pos/repository"
	"github.com/mashingan/smapping"
	"math/rand"
	"strconv"
)

type ProductService interface {
	Insert(p dto.CreateProductDTO) entity.Product
	Update(p dto.UpdateProductDTO) entity.Product
	Delete(productID uint32, article entity.Product)
	All() []entity.Product
	FindByID(productID uint32) entity.Product
	IsAllowedToEdit(merchantID string, productID uint32) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepo}
}

func (service *productService) Insert(p dto.CreateProductDTO) entity.Product {
	mapped := smapping.MapFields(&p)
	productToCreate := entity.Product{}
	err := smapping.FillStruct(&productToCreate, mapped)
	if err != nil {
		panic(err)
	}
	//fmt.Println("image: ", productToCreate.Image)
	productToCreate.Sku = strconv.FormatInt(rand.Int63n(1e16), 10)
	//fmt.Println("sku: ", productToCreate.Sku)
	res := service.productRepository.Insert(productToCreate)
	return res
}

func (service *productService) Update(p dto.UpdateProductDTO) entity.Product {
	mapped := smapping.MapFields(&p)
	productToUpdate := entity.Product{}
	err := smapping.FillStruct(&productToUpdate, mapped)
	if err != nil {
		panic(err)
	}

	res := service.productRepository.Update(productToUpdate)
	return res
}

func (service *productService) Delete(productID uint32, product entity.Product) {
	service.productRepository.Delete(productID, product)
}

func (service *productService) All() []entity.Product {
	return service.productRepository.All()
}

func (service *productService) FindByID(productID uint32) entity.Product {
	return service.productRepository.FindByID(productID)
}

func (service *productService) IsAllowedToEdit(merchantID string, productID uint32) bool {
	p := service.productRepository.FindByID(productID)
	_merchantID := fmt.Sprintf("%v", p.MerchantID)

	return merchantID == _merchantID
}
