package controller

import (
	"net/http"

	logger "github.com/danzok/goCrud.git/src/configuration/logs"
	"github.com/danzok/goCrud.git/src/configuration/validation"
	"github.com/danzok/goCrud.git/src/controller/model/request"
	"github.com/danzok/goCrud.git/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidationUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   int8(userRequest.Age),
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
