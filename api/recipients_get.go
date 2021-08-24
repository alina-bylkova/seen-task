package api

import (
	"net/http"

	"github.com/alinabylkova/seen-task/db"
	"github.com/gin-gonic/gin"
)

func RecipientsGet(db db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, _ := db.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
