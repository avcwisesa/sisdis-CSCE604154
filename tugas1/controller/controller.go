package controller

import (
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
	"context"

	d "github.com/avcwisesa/sisdis/tugas1/database"
	m "github.com/avcwisesa/sisdis/tugas1/model"
	comm "github.com/avcwisesa/sisdis/tugas1/communicator"
)

type Controller interface {
	Register(context.Context, m.Customer) (m.Customer, error)
	GetCustomer(context.Context, string) (m.Customer, error)
	GetTotalSaldo(context.Context, string) (int, error)
	Transfer(context.Context, string, uint) (int, error)
	TransferMinus(context.Context, string, uint) (int, error)
}

type controller struct {
	database    d.Database
	host       string
}

func New(host string, database d.Database) Controller {
	return &controller{
		database:    database,
		host: host,
	}
}

func (c *controller) Register(ctx context.Context, customer m.Customer) (m.Customer, error) {
	select {
	case <-ctx.Done():
		return m.Customer{}, ctx.Err()
	default:
	}

	_, err := c.database.CreateCustomer(customer)
	if err != nil {
		log.Println(err)
		return m.Customer{}, err
	}

	return customer, nil
}

func (c *controller) GetCustomer(ctx context.Context, userID string) (m.Customer, error) {
	select {
	case <-ctx.Done():
		return m.Customer{}, ctx.Err()
	default:
	}

	customer, err := c.database.GetCustomerByID(userID)
	if err != nil {
		log.Println(err)
		return m.Customer{}, err
	}

	return customer, nil
}

func (c *controller) GetTotalSaldo(ctx context.Context, userID string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	jsonFile, err := os.Open("quorum.json")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer jsonFile.Close()

	buf, _ := ioutil.ReadAll(jsonFile)

	var quorum map[string]string
	json.Unmarshal(buf, &quorum)

	var saldo int

	if userID == c.host {

		customer, err := c.GetCustomer(ctx, userID)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		saldo = int(customer.Balance)

		for id, host := range quorum {
			if id == c.host {
				continue
			}
			tmpSaldo, err := comm.GetSaldo(host, id)
			if err != nil {
				log.Println(err)
				return 0, err
			}

			saldo += tmpSaldo
		}

	} else {

		saldo, err = comm.GetTotalSaldo(quorum[userID], userID)
		if err != nil {
			log.Println(err)
			return 0, err
		}

	}

	return saldo, nil
}

func (c *controller) Transfer(ctx context.Context, userID string, nilai uint) (int, error) {
	select {
	case <-ctx.Done():
		return -99, ctx.Err()
	default:
	}

	customer, err := c.database.GetCustomerByID(userID)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	customer.Balance = customer.Balance + nilai

	customer, err = c.database.UpdateCustomer(customer)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return 1, nil
}

func (c *controller) TransferMinus(ctx context.Context, userID string, nilai uint) (int, error) {
	select {
	case <-ctx.Done():
		return -99, ctx.Err()
	default:
	}

	customer, err := c.database.GetCustomerByID(userID)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	customer.Balance = customer.Balance - nilai

	customer, err = c.database.UpdateCustomer(customer)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return 1, nil
}
