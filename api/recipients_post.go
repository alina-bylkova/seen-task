package api

import (
	"net/http"

	"github.com/alinabylkova/seen-task/db"
	"github.com/alinabylkova/seen-task/model"
	"github.com/gin-gonic/gin"
)

// PostRecipient created a new recipient in the database based on the provided name, email and phone number
func PostRecipient(dbLayer db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &model.Recipient{}
		if err := c.ShouldBindJSON(requestBody); isErrorCaught(err, c) {
			return
		}
		if err := validateEmail(requestBody.Email); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		result, err := dbLayer.Add(requestBody)
		if isErrorCaught(err, c) {
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}
