package routes

import (
	"consumer-golang/controllers"
	"github.com/gin-gonic/gin"
)

type ConsumerRouteController struct {
	consumerController controllers.ConsumerController
}

func NewRouteConsumerController(consumerController controllers.ConsumerController) ConsumerRouteController {
	return ConsumerRouteController{consumerController}
}

func (pc *ConsumerRouteController) ConsumerRoute(rg *gin.RouterGroup) {
	router := rg.Group("consumers")
	router.POST("/", pc.consumerController.CreateConsumer)
	router.GET("/", pc.consumerController.FindConsumers)
	router.PUT("/:consumerId", pc.consumerController.UpdateConsumer)
	router.GET("/:consumerId", pc.consumerController.FindConsumerById)
	router.DELETE("/:consumerId", pc.consumerController.DeleteConsumer)
}
