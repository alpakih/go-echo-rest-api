package service

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"go-echo-rest-api/app/http/request"
	"go-echo-rest-api/app/models"
	"go-echo-rest-api/app/models/mapper"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		db,
	}
}

func (r *CustomerService) FindByEmail(email string) (c models.Customer, err error) {

	if err := r.db.Where("email = ?", email).First(&c).Error; err != nil {
		log.Debug("ERROR", err)
		return c, err
	}
	return c, err
}

func (r *CustomerService) FindById(id string) (c models.Customer, err error) {

	if err := r.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Debug("ERROR", err)
		return c, err
	}
	return c, err
}

// FindAll
func (r *CustomerService) FindAll() (interface{}, error) {
	var customer []models.Customer
	err := r.db.Find(&customer).Error

	var length = len(customer)
	serialized := make([]mapper.CustomerMapper, length, length)

	for k, _ := range customer {
		serialized[k] = customer[k].List()
	}

	if err != nil {
		log.Debug("ERROR", err)
		return serialized, err
	}
	return serialized, nil
}

// Insert
func (r *CustomerService) Insert(customerRequest request.CustomerRequest) (customer models.Customer, err error) {
	customer.Name = customerRequest.Name
	customer.Email = customerRequest.Email
	customer.Phone = customerRequest.Phone
	customer.Password, err = customer.HashPassword(customerRequest.Password)
	customer.Address = customerRequest.Address
	err = r.db.Create(&customer).Error
	if err != nil {
		log.Debug("ERROR", err)
		return customer, err
	}
	return customer, nil
}

// Update
func (r *CustomerService) Update(customerUpdateRequest request.CustomerUpdateRequest, id string) (customer models.Customer, err error) {

	data, err := r.FindById(id)
	if err != nil {
		return customer, err
	}
	data.Name = customerUpdateRequest.Name
	data.Email = customerUpdateRequest.Email
	data.Phone = customerUpdateRequest.Phone
	data.Address = customerUpdateRequest.Address

	if err := r.db.Save(&customer).Error; err != nil {
		log.Debug("ERROR", err)

		return customer, err
	}

	return customer, nil
}

// Delete
func (r *CustomerService) Destroy(id string) error {
	return r.db.Delete(&models.Customer{ID: id}).Error
}
