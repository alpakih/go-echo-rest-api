package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"go-echo-rest-api/app/models/mapper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Customer struct {
	ID        string     `gorm:"column:id;primary_key:true"`
	Name      string     `gorm:"type:varchar(100);column:name"`
	Email     string     `gorm:"type:varchar(50);column:email;unique_index;not null"`
	Phone     string     `gorm:"type:varchar(13);column:phone;unique_index;not null"`
	Password  string     `gorm:"type:varchar(255);column:password;not null"`
	Address   string     `gorm:"column:address;json:address"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;sql:index"`
}

func (c *Customer) Map() interface{} {
	customerMapper := mapper.CustomerMapper{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		Address: c.Address,
	}
	return customerMapper
}

// Serialize serializes sms data
func (c Customer) List() mapper.CustomerMapper {
	return mapper.CustomerMapper{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		Address: c.Address,
	}
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Customer) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("id", uuid.New().String()); err != nil {
		log.Fatal("Error UUID Generate")
	}
	return nil
}

func (c *Customer) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (c *Customer) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(plain))
	return err == nil
}