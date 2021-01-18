package controllers

import (
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
)

// Get Content of Diary
func GetDiary(c *gin.Context) {
	year := c.Param("year")
	quarter := c.Param("quarter")

	token, err := helpers.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, &models.Response{Message: err.Error()})
		return
	}

	uuid, err := db.FetchUUID(token.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, &models.Response{Message: "Your session was expired"})
		return
	}
	if uuid != token.UUID {
		c.JSON(http.StatusForbidden, &models.Response{Message: "Your session invalid"})
		return
	}

	db := db.Connect()
	defer db.Close()

	// Get the list of diary from mysql
	rows, err := db.Query("SELECT content, date FROM diary WHERE YEAR(date)=? AND QUARTER(date)=?", year, quarter)
	if err != nil {
		c.JSON(http.StatusNotFound, &models.Response{Message: "There is no diary"})
		return
	}

	var res = []models.DiaryContent{}
	var d models.DiaryContent
	for rows.Next() {
		if err := rows.Scan(&d.Content, &d.Date); err != nil {
			c.JSON(http.StatusInternalServerError, &models.Response{Message: "Internal Server Error"})
			return
		}
		res = append(res, d)
	}

	c.JSON(http.StatusOK, res)
}
