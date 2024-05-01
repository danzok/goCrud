package service

import (
	"fmt"

	logger "github.com/danzok/goCrud.git/src/configuration/logs"
	"github.com/danzok/goCrud.git/src/configuration/rest_err"
	"github.com/danzok/goCrud.git/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser Model", zap.String("jorney", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(ud)
	return nil
}
