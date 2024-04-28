package controller

import (
	"fmt"

	"github.com/danzok/goCrud.git/src/configuration/rest_err"
	"github.com/danzok/goCrud.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect filds, error=%s\n", err.Error()),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
