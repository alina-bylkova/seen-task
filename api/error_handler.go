package api

import (
	"errors"
	"net/http"

	"github.com/alinabylkova/seen-task/db"
	"github.com/gin-gonic/gin"
)

const errorMsg string = "Internal error"

func isErrorCaught(err error, c *gin.Context) bool {
	if err != nil {
		if errors.Is(err, &db.DbError{}) {
			c.JSON(http.StatusInternalServerError, errorMsg)
			return true
		}
		c.JSON(http.StatusNotFound, err.Error())
		return true
	}
	return false
}
