package main

import (
	"fmt"

	"example.com/go-guid/controllers"
	"example.com/go-guid/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	db := database.DatabaseConnection()
	r := gin.Default()
	r = controllers.ProductRouter(r, db)
	r.Run(":5000")
}
