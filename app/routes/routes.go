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
		v1.GET("/thoi-gian-bao-hanh", controllers.FilterWarrantyTime)
		v1.POST("/thoi-gian-bao-hanh", controllers.CreateWarrantyTime)
		v1.PUT("/thoi-gian-bao-hanh", controllers.UpdateWarrantyTime)
		v1.DELETE("/thoi-gian-bao-hanh/:id", controllers.DeleteWarrantyTime)

		//loai giam gia
		v1.GET("/loai-giam-gia", controllers.FilterDiscountType)
		v1.POST("/loai-giam-gia", controllers.CreateDiscountType)
		v1.PUT("/loai-giam-gia", controllers.UpdateDiscountType)
		v1.DELETE("/loai-giam-gia/:id", controllers.DeleteDiscountType)

		//don vi tinh
		v1.GET("/don-vi-tinh", controllers.FilterUnit)
		v1.POST("/don-vi-tinh", controllers.CreateUnit)
		v1.PUT("/don-vi-tinh", controllers.UpdateUnit)
		v1.DELETE("/don-vi-tinh/:id", controllers.DeleteUnit)

		//loai san pham
		v1.GET("/loai-san-pham", controllers.FilterProductType)
		v1.POST("/loai-san-pham", controllers.CreateProductType)
		v1.PUT("/loai-san-pham", controllers.UpdateProductType)
		v1.DELETE("/loai-san-pham/:id", controllers.DeleteProductType)

		//san pham
		v1.GET("/san-pham", controllers.FilterProduct)
		v1.POST("/san-pham", controllers.CreateProduct)
		v1.PUT("/san-pham", controllers.UpdateProduct)
		v1.DELETE("/san-pham/:id", controllers.DeleteProduct)

		//nhan vien
		v1.GET("/nhan-vien", controllers.FilterEmployee)
		v1.POST("/nhan-vien", controllers.CreateEmployee)
		v1.PUT("/nhan-vien", controllers.UpdateEmployee)
		v1.DELETE("/nhan-vien/:id", controllers.DeleteEmployee)

		//chuc vu
		v1.GET("/chuc-vu", controllers.FilterRole)
		v1.POST("/chuc-vu", controllers.CreateRole)
		v1.PUT("/chuc-vu", controllers.UpdateRole)
		v1.DELETE("/chuc-vu/:id", controllers.DeleteRole)

		//khach hang
		v1.GET("/khach-hang", controllers.FilterCustomer)
		v1.POST("/khach-hang", controllers.CreateCustomer)
		v1.PUT("/khach-hang", controllers.UpdateCustomer)
		v1.DELETE("/khach-hang/:id", controllers.DeleteCustomer)

		//kho
		v1.GET("/kho", controllers.FilterWareHouse)
		v1.POST("/kho", controllers.CreateWareHouse)
		v1.PUT("/kho", controllers.UpdateWareHouse)
		v1.DELETE("/kho/:id", controllers.DeleteWareHouse)

		//quyen
		v1.GET("/quyen", controllers.GetPermission)
		v1.PATCH("/quyen", controllers.ModifyPermission)

		//hoa don nhap kho
		v1.GET("/hoa-don-nhap-kho", controllers.FilterImportInvoice)
		v1.POST("/hoa-don-nhap-kho", controllers.CreateImportInvoice)

		//chi tiet san pham
		v1.GET("/chi-tiet-san-pham/:product_id", controllers.GetProductDetail)

		//dang nhap
		v1.GET("/dang-nhap", controllers.Login)
	}

	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}