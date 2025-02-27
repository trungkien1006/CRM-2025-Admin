package controllers

import (
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
// @Router /thoi_gian_bao_hanh [get]
func FilterWarrantyTime(c *gin.Context){
	var req requests.Filter
	var res responses.Filter[db.Thoi_gian_bao_hanh]

	if err := Filter(&req, &res, c, "thoi_gian_bao_hanh"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"message": "lay thoi gian bao hanh thanh cong",
	})
}

// @Summary Create Warranty Time
// @Description Create a new warranty time entry
// @Tags warranty
// @Accept application/json
// @Produce json
// @Param warranty_time body requests.Thoi_gian_bao_hanh_create true "Warranty time data"
// @Router /thoi_gian_bao_hanh/create [post]
func CreateWarrantyTime(c *gin.Context) {
	var req requests.Thoi_gian_bao_hanh_create
	var res responses.Thoi_gian_bao_hanh_create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateWarrantyTimeExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"message": "them thoi gian bao hanh thanh cong",
	})
}

// @Summary Update Warranty Time
// @Description Update an existing warranty time entry
// @Tags warranty
// @Accept application/json
// @Produce json
// @Param warranty_time body requests.Thoi_gian_bao_hanh_update true "Updated warranty time data"
// @Router /thoi_gian_bao_hanh/update [put]
func UpdateWarrantyTime(c *gin.Context) {
	var req requests.Thoi_gian_bao_hanh_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateWarrantyTimeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat thoi gian bao hanh thanh cong",
	})
}

// @Summary Delete Warranty Time
// @Description Delete an existing warranty time entry
// @Tags warranty
// @Accept application/json
// @Produce json
// @Param id path string true "Warranty time ID to be deleted"
// @Router /thoi_gian_bao_hanh/delete [delete]
func DeleteWarrantyTime(c *gin.Context) {
	var req requests.Thoi_gian_bao_hanh_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteWarrantyTimeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat thoi gian bao hanh thanh cong",
	})
}