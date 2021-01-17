package controllers

import (
	"errors"
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
)

// logout
func Logout(c *gin.Context) {
	token := &models.TokenDetails{}
	token, err := helpers.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]string{"error_message": err.Error()})
		return
	}

	err = deleteRedisToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Success Logout"})
}

// Delete token from Redis
func deleteRedisToken(token *models.TokenDetails) error {
	uuid, err := db.FetchUUID(token.Username)
	if err != nil {
		return errors.New("You're session was expired")
	}
	if uuid != token.UUID {
		return errors.New("You're session invalid")
	}

	result, err := db.Redis().Del(token.Username).Result()
	if err != nil {
		return err
	}

	if result != 1 {
		return errors.New("Something went wrong")
	}
	return nil
}
