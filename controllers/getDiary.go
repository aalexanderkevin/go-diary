package controllers

import (
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
)

// Get Diary godoc
// @Security bearerAuth
// @Summary Diary services
// @Description Get Content of the Diary
// @Tags diary
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Param quarter path int true "Quarter"
// @Success 200 {object} models.Diary
// @Failure 401 {object} models.ErrorResponse
// @Router /diary/{year}/{quarter} [get]
func GetDiary(c *gin.Context) {
	year := c.Param("year")
	quarter := c.Param("quarter")

	token, err := helpers.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, &models.ErrorResponse{Message: err.Error()})
		return
	}

	uuid, err := db.FetchUUID(token.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, &models.ErrorResponse{Message: "Your session was expired"})
		return
	}
	if uuid != token.UUID {
		c.JSON(http.StatusForbidden, &models.ErrorResponse{Message: "Your session invalid"})
		return
	}

	db := db.Connect()
	defer db.Close()

	// Get the list of diary from mysql
	rows, err := db.Query("SELECT content, date FROM diary WHERE YEAR(date)=? AND QUARTER(date)=?", year, quarter)
	if err != nil {
		c.JSON(http.StatusNotFound, &models.ErrorResponse{Message: "There is no diary"})
		return
	}

	var res = []models.Diary{}
	var d models.Diary
	for rows.Next() {
		if err := rows.Scan(&d.Content, &d.Date); err != nil {
			c.JSON(http.StatusInternalServerError, &models.ErrorResponse{Message: "Internal Server Error"})
			return
		}
		res = append(res, d)
	}

	c.JSON(http.StatusOK, res)
}
