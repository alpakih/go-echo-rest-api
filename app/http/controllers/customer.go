package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-echo-rest-api/app/helpers"
	"go-echo-rest-api/app/http/request"
	"go-echo-rest-api/app/http/response"
	"go-echo-rest-api/app/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type CustomerController struct {
	db *gorm.DB
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		db,
	}
}

func (cc *CustomerController) Login(c echo.Context) error {
	customer := request.CustomerLoginRequest{}
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(customer); err != nil {
		errorData := make(echo.Map)
		for _, v := range err.(validator.ValidationErrors) {
			errorData[v.Field()] = v.Tag()
		}
		return c.JSON(http.StatusBadRequest, response.ValidationError(errorData))
	}
	data, err := service.NewCustomerService(cc.db).FindByEmail(customer.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.NewError(err))
	}
	if !data.CheckPassword(customer.Password) {
		return c.JSON(http.StatusForbidden, helpers.AccessForbidden())
	}
	return c.JSON(http.StatusOK, response.TokenResponse(&data))
}

func (cc *CustomerController) GetByID(c echo.Context) error {
	id := c.Param("id")
	data, err := service.NewCustomerService(cc.db).FindById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.NotFound())
	}
	return c.JSON(http.StatusOK, response.Single(&data))
}

func (cc *CustomerController) Store(c echo.Context) error {
	customer := request.CustomerRequest{}
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(customer); err != nil {
		errorData := make(echo.Map)
		for _, v := range err.(validator.ValidationErrors) {
			errorData[v.Field()] = v.Tag()
		}
		return c.JSON(http.StatusBadRequest, response.ValidationError(errorData))
	}

	data, err := service.NewCustomerService(cc.db).Insert(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequest())
	}

	return c.JSON(http.StatusOK, response.Single(&data))
}

func (cc *CustomerController) FindAll(c echo.Context) error {
	data, err := service.NewCustomerService(cc.db).FindAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, response.BadRequest())
	}
	return c.JSON(http.StatusOK, response.List(&data))
}

func (cc *CustomerController) Update(c echo.Context) error {
	id := c.Param("id")
	customer := request.CustomerUpdateRequest{}

	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(customer); err != nil {
		errorData := make(echo.Map)
		for _, v := range err.(validator.ValidationErrors) {
			errorData[v.Field()] = v.Tag()
		}
		return c.JSON(http.StatusBadRequest, response.ValidationError(errorData))
	}

	data, err := service.NewCustomerService(cc.db).Update(customer, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response.Single(&data))
}

func (cc *CustomerController) Destroy(c echo.Context) error {

	id := c.Param("id")
	if err := service.NewCustomerService(cc.db).Destroy(id); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response.Success("Success Delete record"))
}

func (cc *CustomerController) CurrentUser(c echo.Context) error {
	key := c.Get("user").(string)

	data, err := service.NewCustomerService(cc.db).FindById(key)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.NotFound())
	}
	return c.JSON(http.StatusOK, response.Single(&data))

}
