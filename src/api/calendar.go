package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bitbucket.org/marcoboschetti/visionmonk/src/service"
)

func PostNewCalendarEvent(c *gin.Context) {
	input := struct {
		CreatorID   string `json:"creator_id"`
		SourceImgID string `json:"source_img_id"`
	}{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newCalendarEvent, err := service.PostNewCalendarEvent(input.CreatorID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newCalendarEvent)
}

func GetCalendarEvent(c *gin.Context) {
	updatedCalendarEvent, err := service.GetCalendarEventsForToday()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated_calendar_event": updatedCalendarEvent})
}
