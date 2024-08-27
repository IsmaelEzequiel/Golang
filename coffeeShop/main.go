package main

import (
	"CoffeeShop/domain"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("APP_PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(ctx *gin.Context) {
		coffee, _ := domain.GetCoffees()
		ctx.HTML(http.StatusOK, "index.html", gin.H{"list": coffee.CoffeeList})
	})

	r.GET("/coffees", func(ctx *gin.Context) {
		coffees, err := domain.GetCoffees()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while getting coffeelist",
				"error":   err,
			})
		}

		ctx.JSON(http.StatusOK, coffees)
	})

	r.GET("/coffees/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		coffee := domain.IsCoffeeAvailable(param)
		ctx.JSON(http.StatusOK, coffee)
	})

	fmt.Printf("Starting the app on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
