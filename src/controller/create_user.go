package controller

import (
	"fmt"
	"log"

	"github.com/danzok/goCrud.git/src/configuration/validation"
	"github.com/danzok/goCrud.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying marshal object, error=%s\n", err.Error())
		restErr := validation.ValidationUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
