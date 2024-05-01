package service

import (
	"github.com/danzok/goCrud.git/src/configuration/rest_err"
	"github.com/danzok/goCrud.git/src/model"
)

func NewUserDomainService() userDomainService {
	return userDomainService{}
}

type userDomainService struct {
}
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
