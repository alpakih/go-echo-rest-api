package response

import (
	"go-echo-rest-api/app/helpers"
	"go-echo-rest-api/app/models"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data, omitempty"`
}

type LoginResponse struct {
	Meta  Meta   `json:"meta"`
	Token string `json:"token"`
}

type CustomerResponse struct {
	Meta     Meta     `json:"meta"`
	Customer *Customer `json:"customer"`
}

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Token   string `json:"token"`
}

func CustomerComposer(u *models.Customer) *Customer {
	r := new(Customer)
	r.Name = u.Name
	r.Email = u.Email
	r.Phone = u.Phone
	r.Address = u.Address
	r.Token = helpers.GenerateJWT(u.ID)
	return r
}
