package controllers

import (
	"database/sql"
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Login service
func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]string{"error_message": "invalid JSON format"})
		return
	}

	var hash string
	username := u.Username
	password := u.Password
	h := helpers.Hash{}
	// Initialize db connection
	db := db.Connect()
	defer db.Close()

	// Get the password from mysql
	row := db.QueryRow("SELECT password, username FROM user WHERE username=? OR email=?", username, username)
	err := row.Scan(&hash, &username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, map[string]string{"error_message": "Invalid username or email"})
			return
		}
		c.JSON(http.StatusInternalServerError, map[string]string{"error_message": "Internal Server Error"})
		return
	}

	// Check the password
	compareError := h.Compare(hash, password)
	if compareError != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error_message": compareError.Error()})
		return
	}

	// Create JWT token
	token, err := helpers.CreateToken(username)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Save token details to Redis
	saveErr := saveUUID(token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	c.JSON(http.StatusOK, map[string]string{"token": token.Token})
}

func saveUUID(token *models.TokenDetails) error {
	exp := time.Unix(token.Exp, 0)
	now := time.Now()

	errAccess := db.Redis().Set(token.Username, token.UUID, exp.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}
