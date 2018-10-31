package database

import (
	"time"

	"github.com/jinzhu/gorm"

	m "github.com/avcwisesa/sisdis/tugas1/model"
)

// database holds the structure for Database
// its encapsulating GORM ORM
type database struct {
	client *gorm.DB
}

// Database holds the contract for DB interface
type Database interface {
	// Migrate should holds implementation for migrating any struct to database table
	Migrate(interface{})
}

// New is used to initiate interface for the DB client
// It returns the interface for the client
func New(client *gorm.DB) Database {
	return &database{client: client}
}

// Migrate is an utility to migrate any golang struct into database table
// Encapsulating GORM Automigrate
func (d *database) Migrate(i interface{}) {
	d.client.AutoMigrate(i)
}

func (d *database) CreateCustomer(customer m.Customer) (m.Customer, error) {

	//set created date
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	// put product to database
	if err := d.client.Where(&m.Customer{UserID: customer.UserID}).First(&m.Customer{}).Error; err != nil {
		d.client.Create(&customer)
	}

	// update product to latest
	d.client.Where(&m.Customer{UserID: customer.UserID}).First(&customer)

	return customer, nil
}

func (d *database) GetCustomerByID(id string) (m.Customer, error) {

	var customer m.Customer
	if err := d.client.First(&customer, id).Error; err != nil {
		return m.Customer{}, err
	}

	return customer, nil
}

func (d *database) UpdateCustomer(customer m.Customer) (m.Customer, error) {

	var customerOld m.Customer
	if err := d.client.First(&customerOld, customer.UserID).Error; err != nil {
		return m.Customer{}, err
	}

	d.client.Model(&customerOld).Updates(customer)

	return customer, nil
}
