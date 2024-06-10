package controllers

import (
	"consumer-golang/models"
	"consumer-golang/request"
	"consumer-golang/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type ConsumerController struct {
	DB *gorm.DB
}

func NewConsumerController(DB *gorm.DB) ConsumerController {
	return ConsumerController{DB}
}

func (pc *ConsumerController) CreateConsumer(ctx *gin.Context) {
	var payload *request.CreateConsumerRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//
	newConsumerEntity := models.Consumer{
		Name:    payload.Name,
		Address: payload.Address,
		Email:   payload.Email,
	}
	//
	result := pc.DB.Create(&newConsumerEntity)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Tên khách hàng không được trùng nhau"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	res := response.ConsumerDTO{
		Id:      newConsumerEntity.ID.String(),
		Name:    newConsumerEntity.Name,
		Address: newConsumerEntity.Address,
		Email:   newConsumerEntity.Email,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (pc *ConsumerController) UpdateConsumer(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	var payload *request.UpdateConsumerRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedConsumerEntity models.Consumer

	result := pc.DB.First(&updatedConsumerEntity, "id = ?", consumerId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Không tồn tại khách hàng để cập nhật trong hệ thống"})
		return
	}

	consumerToUpdate := models.Consumer{
		Name:    payload.Name,
		Address: payload.Address,
		Email:   payload.Email,
	}

	pc.DB.Model(&updatedConsumerEntity).Updates(consumerToUpdate)
	//
	var res response.ConsumerDTO = response.ConsumerDTO{
		Id:      updatedConsumerEntity.ID.String(),
		Name:    updatedConsumerEntity.Name,
		Address: updatedConsumerEntity.Address,
		Email:   updatedConsumerEntity.Email,
	}
	//

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (pc *ConsumerController) FindConsumerById(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")
	var consumerEntity models.Consumer
	result := pc.DB.First(&consumerEntity, "id = ?", consumerId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Không tồn tại khách hàng này trong hệ thống"})
		return
	}
	//
	var res response.ConsumerDTO = response.ConsumerDTO{
		Id:      consumerEntity.ID.String(),
		Name:    consumerEntity.Name,
		Address: consumerEntity.Address,
		Email:   consumerEntity.Email,
	}
	//
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (pc *ConsumerController) FindConsumers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var consumers []models.Consumer
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&consumers)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	//
	var listDTO []response.ConsumerDTO
	for _, consumer := range consumers {
		var dtoConsumer response.ConsumerDTO = response.ConsumerDTO{
			Id:      consumer.ID.String(),
			Name:    consumer.Name,
			Address: consumer.Address,
			Email:   consumer.Email,
		}
		listDTO = append(listDTO, dtoConsumer)
	}
	//
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(listDTO), "data": listDTO})
}

func (pc *ConsumerController) DeleteConsumer(ctx *gin.Context) {
	consumerId := ctx.Param("consumerId")

	result := pc.DB.Delete(&models.Consumer{}, "id = ?", consumerId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Không tồn tại khách hàng để xóa trong hệ thống"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Xóa khách hàng với mã định danh " + consumerId + " thành công"})
}
