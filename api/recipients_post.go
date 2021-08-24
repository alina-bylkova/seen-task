package api

import (
	"net/http"

	"github.com/alinabylkova/seen-task/db"
	"github.com/gin-gonic/gin"
)

func RecipientsPost(db db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusCreated, "")

	}
}
