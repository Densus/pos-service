package controller

import (
	"fmt"
	"github.com/densus/pos_service/helper"
	"github.com/densus/pos_service/pos/model/dto"
	"github.com/densus/pos_service/pos/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type OutletProductController interface {
	SetPrice(ctx *gin.Context)
}

type outletProductController struct {
	outletProductService service.OutletProductService
	jwtService     service.JWTService
}

func NewOutletProductController(productService service.OutletProductService, jwtServ service.JWTService) OutletProductController {
	return &outletProductController{outletProductService: productService, jwtService: jwtServ}
}

func (o *outletProductController) SetPrice(ctx *gin.Context) {
	var setPriceDTO dto.SetPriceDTO
	errDTO := ctx.ShouldBind(&setPriceDTO)
	if errDTO != nil {
		res := helper.ErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New()
	//validate input client based on tag validate model dto.UpdateUserDTO
	if errDTO := validate.Struct(setPriceDTO); errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	_idProduct := ctx.Param("product_id")
	productId, _ := strconv.ParseUint(_idProduct, 0, 32)
	setPriceDTO.ProductID = uint32(productId)

	_idOutlet := ctx.Param("outlet_id")
	outletId, _ := strconv.ParseUint(_idOutlet, 0, 32)
	setPriceDTO.OutletID = uint32(outletId)

	fmt.Println("price: ", setPriceDTO)

	u := o.outletProductService.SetPrice(setPriceDTO.ProductID, setPriceDTO.OutletID, setPriceDTO)
	res := helper.SuccessResponse(true, "OK", u)
	ctx.JSON(http.StatusOK, res)
}
