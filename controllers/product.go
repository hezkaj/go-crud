package controllers

import (
	"database/sql"

	//"fmt"
	"net/http"

	"example.com/go-guid/database"
	"example.com/go-guid/services"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.Engine, db *sql.DB) *gin.Engine {
	r = findOneProduct(r, db)
	r = findAllProduct(r, db)
	r = createProduct(r, db)
	r = updateProduct(r, db)
	r = updateCheckBoxProduct(r, db)
	r = deleteProduct(r, db)
	return r
}

func createProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.POST("/products", func(c *gin.Context) {
		var prod *database.Product
		if err := c.ShouldBind(&prod); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		services.CreateProduct(prod, DB)
		c.JSON(http.StatusOK, gin.H{
			"message": "product created successfully",
		})
	})
	return r
}
func findOneProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, _ := services.FindOneProduct(DB, &id)
		if product == nil {
			c.JSON(http.StatusNotFound, product)
			return
		}
		c.JSON(http.StatusOK, product)
	})
	return r
}
func findAllProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.GET("/products", func(c *gin.Context) {
		products, _ := services.FindAllProduct(DB)
		c.JSON(http.StatusOK, products)
	})
	return r
}
func updateProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.PUT("/products/:id", func(c *gin.Context) {
		var prod database.Product
		id := c.Param("id")
		if err := c.ShouldBind(&prod); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		services.UpdateProduct(&prod, DB, &id)
		product, _ := services.FindOneProduct(DB, &id)
		if product == nil {
			c.JSON(http.StatusNotFound, product)
			return
		}
		c.JSON(http.StatusOK, product)
	})
	return r
}
func updateCheckBoxProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.PATCH("/products/:id", func(c *gin.Context) {
		var prod database.Product
		id := c.Param("id")
		if err := c.ShouldBind(&prod); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		services.UpdateCheckBoxProduct(&prod, DB, &id)
		product, _ := services.FindOneProduct(DB, &id)
		if product == nil {
			c.JSON(http.StatusNotFound, product)
			return
		}
		c.JSON(http.StatusOK, product)
	})
	return r
}
func deleteProduct(r *gin.Engine, DB *sql.DB) *gin.Engine {
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, _ := services.FindOneProduct(DB, &id)
		if product == nil {
			c.JSON(http.StatusNotFound, product)
			return
		}
		services.DeleteProduct(DB, &id)
		c.JSON(http.StatusOK, gin.H{
			"message": "product deleted successfully",
		})
	})
	return r
}
