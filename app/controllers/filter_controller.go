package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"

	"github.com/gin-gonic/gin"
)

func Filter[T any](req *requests.Filter, res *responses.Filter[T], c *gin.Context, tableName string) error {
	if err := c.ShouldBindQuery(&req); err != nil {
		return err
	}

	if err := dao.FilterExec(req, res, tableName); err != nil {
		return err
	}

	return nil
}