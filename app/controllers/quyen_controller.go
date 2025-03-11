package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get Permission
// @Description Get permission by role id
// @Tags permission
// @Accept application/x-www-form-urlencoded
// @Param Chuc_vu_id query int true "Role id"
// @Router /permission [get]
func GetPermission(c *gin.Context) {
	var req requests.Quyen_read
	var res responses.Quyen_read

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.GetPermissionExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"message": "lay quyen thanh cong",
	})
}

// @Summary Get Permission
// @Description Modify permission by Role
// @Tags permission
// @Accept application/json
// @Param Modify body requests.Quyen_modify true "Modify permission"
// @Router /permission [patch]
func ModifyPermission(c *gin.Context) {
	var req requests.Quyen_modify
	var ds_code_quyen []string
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := dao.ModifyPermissionExec(&req, &ds_code_quyen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	quyenJson, err := json.Marshal(ds_code_quyen)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := helpers.Redis.Set(helpers.Ctx, "role:" + strconv.Itoa(int(req.Chuc_vu_id)), quyenJson, 0).Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chinh sua quyen thanh cong",
	})
}