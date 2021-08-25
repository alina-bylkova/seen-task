package api

import (
	"net/http"

	"github.com/alinabylkova/seen-task/db"
	"github.com/alinabylkova/seen-task/model"
	"github.com/alinabylkova/seen-task/model/dto"
	"github.com/gin-gonic/gin"
)

// PostEvent creates or updates event in the database based on the provided recipient_id, video_id and event_type
func PostEvent(dbLayer db.Layer) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &dto.Event{}
		if err := c.ShouldBindJSON(requestBody); isErrorCaught(err, c) {
			return
		}
		if err := validateId(requestBody.RecipientID); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := validateId(requestBody.VideoID); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := requestBody.ValidateEventType(); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		recipient, err := dbLayer.Get(&model.Recipient{ID: requestBody.RecipientID})
		if isErrorCaught(err, c) {
			return
		}
		if len(recipient) == 0 {
			c.JSON(http.StatusNotFound, "Recipient not found")
			return
		}
		if err := dbLayer.AddEvent(requestBody); isErrorCaught(err, c) {
			return
		}
		c.Status(http.StatusCreated)
	}
}
