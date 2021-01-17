package controllers

import (
	"fmt"
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
)

// Create Diary
func CreateDiary(c *gin.Context) {
	var d models.Diary

	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{"error_message": "Invalid JSON format"})
		return
	}

	token, err := helpers.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]string{"error_message": err.Error()})
		return
	}

	uuid, err := db.FetchUUID(token.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]string{"error_message": "Your session was expired"})
		return
	}
	if uuid != token.UUID {
		c.JSON(http.StatusForbidden, map[string]string{"error_message": "Your session invalid"})
		return
	}

	db := db.Connect()
	defer db.Close()

	_, err = db.Exec("INSERT INTO diary (user_id, content, date) values (?,?,?)",
		token.Username,
		d.Content,
		d.Date,
	)
	if err != nil {
		fmt.Println("masuk ke UPDATE")
		_, err = db.Exec("UPDATE diary SET content=? WHERE user_id=? AND date=?",
			d.Content,
			token.Username,
			d.Date,
		)
		if err != nil {
			c.JSON(http.StatusBadGateway, map[string]string{"error_message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Diary added"})
}
