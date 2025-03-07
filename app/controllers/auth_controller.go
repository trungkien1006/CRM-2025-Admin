package controllers

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description Login API
// @Tags auth
// @Accept  json
// @Produce  json
// @Param Login_data body requests.Dang_nhap true "Login data include username, password"
// @Success 200 {object} responses.Dang_nhap
// @Failure 400 {object} map[string]interface{}
// @Router /dang-nhap [post]
func Login(c *gin.Context) {
	var req requests.Dang_nhap
	var res responses.Dang_nhap

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding",
			"message": err.Error(),
		})

		return
	}

	if id, role_id, err := dao.LoginExec(&req, &res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	} else {
		isExist, redisErr := helpers.Redis.Exists(helpers.Ctx, "user:" + strconv.Itoa(int(id))).Result()

		if redisErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "redis check exist",
				"message": redisErr.Error(),
			})

			return
		}

		if isExist == 0 {
			if e := helpers.Redis.Set(helpers.Ctx, "user:" + strconv.Itoa(int(id)), strconv.Itoa(role_id), 0).Err(); e != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "redis save",
					"message": e.Error(),
				})

				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "dang nhap thanh cong",
	})
}