package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Jsondata struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
}

func RoomDHT11(c *gin.Context) {
	reqid := c.MustGet("RequestId").(string)
	var jsoninput Jsondata
	c.BindJSON(&jsoninput)
	fmt.Println(reqid, "data > ", jsoninput.Temperature, jsoninput.Humidity)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
