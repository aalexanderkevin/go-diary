package controllers

import (
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// Insert user credentials to db
func InsertUser(c *gin.Context) {
	var u models.User
	h := helpers.Hash{}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &models.Response{Message: "Invalid JSON format"})
		return
	}

	db := db.Connect()
	defer db.Close()

	username := u.Username
	password := u.Password
	email := u.Email
	name := u.Name
	birthday := u.Birthday

	generatedHash, generateError := h.Generate(password)
	if generateError != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Message: generateError.Error()})
		return
	}

	if err := helpers.VerifyPassword(password); err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{Message: err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO user (username, password, email, name, birthday) values (?,?,?,?,?)",
		username,
		generatedHash,
		email,
		name,
		birthday,
	)

	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				c.JSON(http.StatusConflict, &models.Response{Message: "username already exist"})
				return
			}
		}
		c.JSON(http.StatusBadGateway, &models.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Success"})
	return
}
