package controller

import (
	"bytes"
	"github.com/densus/pos_service/helper"
	"github.com/densus/pos_service/pos/model/dto"
	"github.com/densus/pos_service/pos/model/entity"
	"github.com/densus/pos_service/pos/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
)

type ProductController interface {
	All(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewArticleController(productServ service.ProductService, jwtServ service.JWTService) ProductController {
	return &productController{productService: productServ, jwtService: jwtServ}
}

func (a *productController) All(ctx *gin.Context) {
	var articles []entity.Product = a.productService.All()
	res := helper.SuccessResponse(true, "OK", articles)
	ctx.JSON(http.StatusOK, res)
}

func (a *productController) Insert(ctx *gin.Context) {
	var productCreateDTO dto.CreateProductDTO

	if errDTO := ctx.ShouldBind(&productCreateDTO); errDTO != nil {
		res := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	validate := validator.New()
	//validate input client based on tag validate model dto.CreateArticleDTO
	if errDTO := validate.Struct(productCreateDTO); errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		file, _, err := ctx.Request.FormFile("image")
		defer file.Close()
		if err != nil {
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return
		}
		productCreateDTO.Image = buf.Bytes()

		result := a.productService.Insert(productCreateDTO)
		response := helper.SuccessResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (a *productController) Update(ctx *gin.Context) {
	var productUpdateDTO dto.UpdateProductDTO
	errDTO := ctx.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.ErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New()
	//validate input client based on tag validate model dto.UpdateArticleDTO
	if errDTO := validate.Struct(productUpdateDTO); errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		file, _, err := ctx.Request.FormFile("image")
		defer file.Close()
		if err != nil {
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return
		}
		productUpdateDTO.Image = buf.Bytes()

		_id := ctx.Param("id")
		productId, _ := strconv.ParseUint(_id, 0, 32)
		productUpdateDTO.ID = uint32(productId)
		//fmt.Println("productID: ", productId)
		//fmt.Println("merchantID: ", strconv.Itoa(int(productUpdateDTO.MerchantID)))

		if a.productService.IsAllowedToEdit(strconv.Itoa(int(productUpdateDTO.MerchantID)), productUpdateDTO.ID) {
			id, errID := strconv.ParseUint(strconv.Itoa(int(productUpdateDTO.MerchantID)), 10, 32)
			if errID == nil {
				productUpdateDTO.MerchantID = uint32(id)
			}
			result := a.productService.Update(productUpdateDTO)
			response := helper.SuccessResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		} else {
			response := helper.ErrorResponse("You don't have permission", "You're not the owner", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
		}
	}
}

func (a *productController) Delete(ctx *gin.Context) {
	var product entity.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.ErrorResponse("Failed to get id", "param id is not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	productID := id

	a.productService.Delete(uint32(productID), product)
	res := helper.SuccessResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (a *productController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.ErrorResponse("Param id is not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	_id := uint32(id)
	var product entity.Product = a.productService.FindByID(_id)
	res := helper.SuccessResponse(true, "OK!", product)
	ctx.JSON(http.StatusOK, res)
}
