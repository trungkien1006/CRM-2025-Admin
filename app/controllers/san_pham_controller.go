package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Unit
// @Description Filter unit based on provided filters
// @Tags unit
// @Accept application/json
// @Produce json
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /san-pham [get]
func FilterProduct(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.San_pham]

	if err := Filter(&req, &res, c, "san_pham"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay san pham thanh cong",
	})
}

// @Summary Create Product
// @Description Create a new product entry
// @Tags product
// @Accept  multipart/form-data
// @Produce json
// @Param Product body requests.San_pham_create true "Product data"
// @Router /san-pham/create [post]
func CreateProduct(c *gin.Context) {
	var req requests.San_pham_create
	var res responses.San_pham_create

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := helpers.StoreFile(req.Hinh_anh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	for _, value := range req.Chi_tiet_san_pham {
		if err := helpers.StoreFile(value.Hinh_anh); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}
	}

	if err := dao.CreateProductExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them san pham thanh cong",
	})
}

// @Summary Update Product
// @Description Update an existing product entry
// @Tags product
// @Accept  multipart/form-data
// @Produce json
// @Param Product body requests.San_pham_update true "Updated Product data"
// @Router /san-pham/update [put]
func UpdateProduct(c *gin.Context) {
	var req requests.San_pham_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := helpers.StoreFile(req.Hinh_anh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	for _, value := range req.Chi_tiet_san_pham {
		if err := helpers.StoreFile(value.Hinh_anh); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}
	}

	if err := dao.UpdateProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat san pham thanh cong",
	})
}

// @Summary Delete Product
// @Description Delete an existing product entry
// @Tags product
// @Accept application/json
// @Produce json
// @Param id path string true "product ID to be deleted"
// @Router /san-pham/delete [delete]
func DeleteProduct(c *gin.Context) {
	var req requests.San_pham_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteProductExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa san pham thanh cong",
	})
}