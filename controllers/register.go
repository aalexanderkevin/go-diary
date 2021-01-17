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
		c.JSON(http.StatusUnprocessableEntity, map[string]string{"error_message": "Invalid JSON format"})
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
		c.JSON(http.StatusBadRequest, map[string]string{"error_message": generateError.Error()})
		return
	}

	if err := helpers.VerifyPassword(password); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error_message": err.Error()})
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
				c.JSON(http.StatusConflict, map[string]string{"error_message": "username already exist"})
				return
			}
		}
		c.JSON(http.StatusBadGateway, map[string]string{"error_message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Success"})
	return
}
