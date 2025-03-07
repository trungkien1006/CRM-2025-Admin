package app

import (
	"admin-v1/app/configs"
	"admin-v1/app/helpers"
	// "admin-v1/app/initRedis"
	"admin-v1/app/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	helpers.GormDB = configs.GormConnection()

	r := routes.InitRoute()

	// initRedis.InitRolePermission()

	return r
}