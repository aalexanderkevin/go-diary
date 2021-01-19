package controllers

import (
	"errors"
	"net/http"
	"tap-talk/db"
	"tap-talk/helpers"
	"tap-talk/models"

	"github.com/gin-gonic/gin"
)

// Logout godoc
// @Security bearerAuth
// @Summary Logout services
// @Description Logout user
// @Tags logout
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 401 {object} models.ErrorResponse
// @Router /logout [post]
func Logout(c *gin.Context) {
	token := &models.TokenDetails{}
	token, err := helpers.ExtractToken(c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, &models.ErrorResponse{Message: err.Error()})
		return
	}

	err = deleteRedisToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &models.Response{Message: "Success Logout"})
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
