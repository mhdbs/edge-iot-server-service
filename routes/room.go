package routes

import (
	"mhdbs/edge-iot-server-service/controller"

	"github.com/gin-gonic/gin"
)

func Room(router *gin.RouterGroup) {
	router.POST("/sensor/dht11", controller.RoomDHT11)
}
