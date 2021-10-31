package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petshop/dto"
	"petshop/storage/user_data"
)

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request dto.ListUserRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"error": nil,
				"data": nil,
			})
			return
		}

		users, err := user_data.FindAll(request.Ids)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"error": err.Error(),
				"data": nil,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"error": "",
			"data": users,
		})
		return
	}
}