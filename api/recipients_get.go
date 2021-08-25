package api

import (
	"net/http"
	"strconv"

	"github.com/alinabylkova/seen-task/db"
	"github.com/alinabylkova/seen-task/model"
	"github.com/gin-gonic/gin"
)

// GetRecipientById returns a list containing one recipient based on provided id
func GetRecipientById(dbLayer db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := validateId(id); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		recipient := &model.Recipient{ID: id}
		result, err := dbLayer.Get(recipient)
		if isErrorCaught(err, c) {
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

// GetRecipients returns a list of recipients based on provided name, email or phone number, otherwise all recipients will be returned
func GetRecipients(dbLayer db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		email := c.Query("email")
		phone := c.Query("phone")
		if len(name) == 0 && len(email) == 0 && len(phone) == 0 {
			result, err := dbLayer.GetAll()
			if isErrorCaught(err, c) {
				return
			}
			c.JSON(http.StatusOK, result)
			return
		}

		if err := validateName(name); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := validateEmail(email); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := validatePhone(phone); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		recipient := &model.Recipient{Name: name, Email: email, Phone: phone}
		result, err := dbLayer.Get(recipient)
		if isErrorCaught(err, c) {
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
