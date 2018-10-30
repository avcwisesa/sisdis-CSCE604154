package controller

import (
	"context"

	d "github.com/avcwisesa/sisdis/tugas1/database"
	m "github.com/avcwisesa/sisdis/tugas1/model"
)

type controller struct {
	database    d.Database
}

func New(database Database) Controller {
	return &controller{
		database:    database,
	}
}

func (c *controller) Register(ctx context.Context, customer m.Customer) (m.Customer, error) {
	select {
	case <-ctx.Request.Context().Done():
		return &m.Customer{}, ctx.Err()
	default:
	}

	_, err := c.database.CreateCustomer(customer)
	if err != nil {
		return &m.Customer{}, err
	}

	return customer, nil
}

func (c *controller) GetCustomer(ctx context.Context, userID uint) (m.Customer, error) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	customer, err := c.database.GetCustomerByID(userID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

