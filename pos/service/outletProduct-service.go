package service

import (
	"fmt"
	"github.com/densus/pos_service/pos/model/dto"
	"github.com/densus/pos_service/pos/model/entity"
	"github.com/densus/pos_service/pos/repository"
	"github.com/mashingan/smapping"
)

type OutletProductService interface {
	SetPrice(ProductID, OutletID uint32, dto dto.SetPriceDTO) entity.OutletProduct
}

type outletProductService struct {
	outletProductRepository repository.OutletProductRepository
}

func NewOutletProductService(productRepository repository.OutletProductRepository) OutletProductService {
	return &outletProductService{outletProductRepository: productRepository}
}

func (o *outletProductService) SetPrice(ProductID, OutletID uint32, dto dto.SetPriceDTO) entity.OutletProduct {
	mapped := smapping.MapFields(&dto)
	priceToCreate := entity.OutletProduct{}
	err := smapping.FillStruct(&priceToCreate, mapped)
	fmt.Println("ptc: ", priceToCreate)
	if err != nil {
		panic(err)
	}

	res := o.outletProductRepository.SetPrice(ProductID, OutletID, priceToCreate)
	return res
}
