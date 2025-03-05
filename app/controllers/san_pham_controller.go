package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Product
// @Description Filter product based on provided filters
// @Tags product
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /san-pham [get]
func FilterProduct(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.San_pham_filter]

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
// @Param ten formData string true "Product Name"
// @Param upc formData string true "UPC Code"
// @Param loai_san_pham_id formData int true "Product Type ID"
// @Param file formData file true "Product Image"
// @Param don_vi_tinh_id formData int true "Unit ID"
// @Param vat formData float32 false "VAT (Optional)"
// @Param mo_ta formData string false "Description (Optional)"
// @Param trang_thai formData int true "Status"
// @Param loai_giam_gia_id formData int false "Discount Type ID (Optional)"
// @Param thoi_gian_bao_hanh_id formData int false "Warranty Time ID (Optional)"
// @Success 200 {object} map[string]interface{} "data: San_pham_create, message: them san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /san-pham [post]
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
// @Param id formData int true "Product ID"
// @Param ten formData string true "Product Name"
// @Param upc formData string true "UPC Code"
// @Param loai_san_pham_id formData int true "Product Type ID"
// @Param file formData file false "Product Image (Optional)"
// @Param don_vi_tinh_id formData int true "Unit ID"
// @Param vat formData float32 false "VAT (Optional)"
// @Param mo_ta formData string false "Description (Optional)"
// @Param trang_thai formData int true "Status"
// @Param loai_giam_gia_id formData int false "Discount Type ID (Optional)"
// @Param thoi_gian_bao_hanh_id formData int false "Warranty Time ID (Optional)"
// @Success 200 {object} map[string]interface{} "message: cap nhat san pham thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /san-pham [put]
func UpdateProduct(c *gin.Context) {
	var req requests.San_pham_update

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if req.Hinh_anh != nil {
		if err := helpers.StoreFile(req.Hinh_anh); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}
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
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "product ID to be deleted"
// @Router /san-pham/{id} [delete]
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