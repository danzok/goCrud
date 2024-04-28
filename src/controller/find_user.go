package controller

import (
	"github.com/danzok/goCrud.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Passe um id para a url")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {

}
