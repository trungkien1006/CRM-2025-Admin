package routes

import (
	"net/http"

	"admin-v1/app/controllers"
	_ "admin-v1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// Thêm header CORS cho mỗi request
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                       // Cho phép tất cả các origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS") // Các phương thức HTTP cho phép
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")            // Các header cho phép

		// Xử lý preflight OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	r.Static("/public/images", "./public/images")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//version 1.0
	v1 := r.Group("/api/v1")
	{
		//thoi gian bao hanh
		v1.POST("/thoi-gian-bao-hanh", controllers.FilterWarrantyTime)
		v1.POST("/thoi-gian-bao-hanh/create", controllers.CreateWarrantyTime)
		v1.PUT("/thoi-gian-bao-hanh/update", controllers.UpdateWarrantyTime)
		v1.DELETE("/thoi-gian-bao-hanh/delete", controllers.DeleteWarrantyTime)

		//loai giam gia
		v1.POST("/loai-giam-gia", controllers.FilterDiscountType)
		v1.POST("/loai-giam-gia/create", controllers.CreateDiscountType)
		v1.PUT("/loai-giam-gia/update", controllers.UpdateDiscountType)
		v1.DELETE("/loai-giam-gia/delete", controllers.DeleteDiscountType)

		//don vi tinh
		v1.POST("/don-vi-tinh", controllers.FilterUnit)
		v1.POST("/don-vi-tinh/create", controllers.CreateUnit)
		v1.PUT("/don-vi-tinh/update", controllers.UpdateUnit)
		v1.DELETE("/don-vi-tinh/delete", controllers.DeleteUnit)

		//loai san pham
		v1.POST("/loai-san-pham", controllers.FilterProductType)
		v1.POST("/loai-san-pham/create", controllers.CreateProductType)
		v1.PUT("/loai-san-pham/update", controllers.UpdateProductType)
		v1.DELETE("/loai-san-pham/delete", controllers.DeleteProductType)

		//san pham
		v1.POST("/san-pham", controllers.FilterProduct)
	}

	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}