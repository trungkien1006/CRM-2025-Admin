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
// @Router /loai-san-pham [get]
func FilterProductType(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[db.Loai_san_pham]

	if err := Filter(&req, &res, c, "loai_san_pham"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay loai san pham thanh cong",
	})
}

// @Summary Create Product Type
// @Description Create a new product type entry
// @Tags product type
// @Accept  multipart/form-data
// @Produce json
// @Param Discount_Type body requests.Loai_san_pham_create true "Product Type data"
// @Router /loai-san-pham/create [post]
func CreateProductType(c *gin.Context) {
	var req requests.Loai_san_pham_create
	var res responses.Loai_san_pham_create

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding",
			"message": err.Error(),
		})

		return
	}

	if err := helpers.StoreFile(req.Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateProductTypeExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them loai san pham thanh cong",
	})
}

// @Summary Update Product Type
// @Description Update an existing product type entry
// @Tags product type
// @Accept  multipart/form-data
// @Produce json
// @Param Product_Type body requests.Loai_san_pham_update true "Updated Product Type data"
// @Router /loai-san-pham/update [put]
func UpdateProductType(c *gin.Context) {
	var req requests.Loai_san_pham_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := helpers.StoreFile(req.Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateProductTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat loai san pham thanh cong",
	})
}

// @Summary Delete Product Type
// @Description Delete an existing product type entry
// @Tags product type
// @Accept application/json
// @Produce json
// @Param id path string true "product type ID to be deleted"
// @Router /loai-san-pham/delete [delete]
func DeleteProductType(c *gin.Context) {
	var req requests.Loai_san_pham_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteProductTypeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa loai san pham thanh cong",
	})
}