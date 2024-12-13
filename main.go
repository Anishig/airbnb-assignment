package main

import (
	"airbnb/controllers"
	"airbnb/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("GO")
	database.ConnectDB()

	r := gin.Default()
	r.GET("/:room_id", controllers.GetRoomMetrics)

	r.Run(":8080")
}
