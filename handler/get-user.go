package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petshop/storage/user_data"
)

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, err := user_data.Find(id)
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
			"data": user,
		})
		return
	}
}

