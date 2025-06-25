package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "hyperchain/api/controllers"
    "hyperchain/api/services"
    "hyperchain/api/config"
)

func main() {
    config.LoadConfig()

    if err := services.InitFabric(); err != nil {
        log.Fatalf("Failed to initialize Fabric SDK: %v", err)
    }

    router := gin.Default()

    router.POST("/record", controllers.AddRecord)
    router.GET("/record/:id", controllers.GetRecord)

	log.Println("API running on http://localhost:8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
