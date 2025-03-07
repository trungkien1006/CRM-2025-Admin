package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Employee
// @Description Filter employee based on provided filters
// @Tags employee
// @Accept application/x-www-form-urlencoded
// @Param filters query string false "Filters in JSON format"
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc/desc)"
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Router /nhan-vien [get]
func FilterEmployee(c *gin.Context) {
	var req requests.Filter
	var res responses.Filter[responses.Nhan_vien_filter]

	if err := Filter(&req, &res, c, "nhan_vien"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "lay nhan vien thanh cong",
	})
}

// @Summary Create Employee
// @Description Create a new employee entry
// @Tags employee
// @Accept  json
// @Produce json
// @Param CreateEmployee body requests.Nhan_vien_create true "Employee Create Data"
// @Success 200 {object} map[string]interface{} "data: Nhan_vien_create, message: them nhan vien thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /nhan-vien [post]
func CreateEmployee(c *gin.Context) {
	var req requests.Nhan_vien_create
	var res responses.Nhan_vien_create

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.CreateEmployeeExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "them nhan vien thanh cong",
	})
}

// @Summary Update Employee
// @Description Update an existing employee entry
// @Tags employee
// @Accept  json
// @Produce json
// @Param UpdateEmployee body requests.Nhan_vien_update true "Employee Update Data"
// @Success 200 {object} map[string]interface{} "message: cap nhat nhan vien thanh cong"
// @Failure 400 {object} map[string]string "message: error message"
// @Router /nhan-vien [put]
func UpdateEmployee(c *gin.Context) {
	var req requests.Nhan_vien_update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.UpdateEmployeeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "cap nhat nhan vien thanh cong",
	})
}


// @Summary Delete Employee
// @Description Delete an existing employee entry
// @Tags employee
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "Employee ID to be deleted"
// @Router /nhan-vien/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	var req requests.Nhan_vien_delete

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.DeleteEmployeeExec(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "xoa nhan vien thanh cong",
	})
}