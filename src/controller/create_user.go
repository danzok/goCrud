package controller

import (
	"net/http"

	logger "github.com/danzok/goCrud.git/src/configuration/logs"
	"github.com/danzok/goCrud.git/src/configuration/validation"
	"github.com/danzok/goCrud.git/src/controller/model/request"
	"github.com/danzok/goCrud.git/src/model"
	"github.com/danzok/goCrud.git/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainINterface model.UserDomainInterface
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
