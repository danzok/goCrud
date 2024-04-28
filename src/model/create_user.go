package model

import (
	"fmt"

	logger "github.com/danzok/goCrud.git/src/configuration/logs"
	"github.com/danzok/goCrud.git/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser Model", zap.String("jorney", "createUser"))
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
