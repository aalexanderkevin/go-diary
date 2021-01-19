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

// Login godoc
// @Summary login services
// @Description Authenticates user and provides a JWT to Authorize API calls
// @Tags login
// @Accept json
// @Produce json
// @Param user body models.Login true "Login"
// @Success 200 {object} models.JWT
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var l models.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &models.ErrorResponse{Message: "invalid JSON format"})
		return
	}

	var hash string
	username := l.Username
	password := l.Password
	h := helpers.Hash{}
	// Initialize db connection
	db := db.Connect()
	defer db.Close()

	// Get the password from mysql
	row := db.QueryRow("SELECT password, username FROM user WHERE username=? OR email=?", username, username)
	err := row.Scan(&hash, &username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Invalid username or email"})
			return
		}
		c.JSON(http.StatusInternalServerError, &models.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	// Check the password
	compareError := h.Compare(hash, password)
	if compareError != nil {
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: compareError.Error()})
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

	c.JSON(http.StatusOK, &models.JWT{Token: token.Token})
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
