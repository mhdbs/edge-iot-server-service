package main

import (
	"log"
	"mhdbs/edge-iot-server-service/config"
	"mhdbs/edge-iot-server-service/routes"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			uuid4 := uuid.NewV4()
			requestID = uuid4.String()
		}
		c.Set("RequestId", requestID)
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

func main() {
	config.LoadConfig()
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.Use(RequestID())
	routes.Room(v1.Group("/room"))
	err := router.Run(config.HttpHost + ":" + config.HttpPort)
	if err != nil {
		log.Panic("couldnt able to start the http server", err)
	}
}
