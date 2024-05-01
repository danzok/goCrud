package view

import (
	"github.com/danzok/goCrud.git/src/controller/model/response"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface,) 
response.UserResponse {
 return response.UserResponse {
	ID : "",
	Email: userDomain.GetEmail(),
	Name : userDomain.GetName(),
	Age : userDomain.GetAge()
 }
}
