package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Import Invoice
// @Description Filter import invoice based on provided filters
// @Tags import invoice
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /hoa-don-nhap-kho [get]
func FilterImportInvoice(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Hoa_don_nhap_kho_filter]

	if err := Filter(&req, &res, c, "hoa_don_nhap_kho"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay hoa don nhap kho thanh cong",
	})
}

// @Summary Create Import Invoice
// @Description Create a new import invoice entry
// @Tags import invoice
// @Accept  json
// @Produce json
// @Param Import_Invoice body requests.Hoa_don_nhap_kho_create true "Import Invoice data"
// @Router /hoa-don-nhap-kho [post]
func CreateImportInvoice(c *gin.Context) {
	var req requests.Hoa_don_nhap_kho_create
	var res responses.Hoa_don_nhap_kho_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateImportInvoice(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them hoa don nhap kho thanh cong",
	})
}