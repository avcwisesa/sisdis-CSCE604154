package handler

import (
	"log"
	"github.com/gin-gonic/gin"

	m "github.com/avcwisesa/sisdis/tugas1/model"
	c "github.com/avcwisesa/sisdis/tugas1/controller"
)

// handler holds the structure for Handler
type handler struct {
	controller c.Controller
	quorum int
}

// Handler holds the contract for Handler
type Handler interface {
	// Ping should handle healthcheck in top level routing
	Ping(*gin.Context)
	// RegisterHandler should handle new customer register
	Register(*gin.Context)
	// GetSaldoHandler should handle get balance
	GetSaldo(*gin.Context)
	// GetTotalSaldoHandler should handle get balance from all branches
	GetTotalSaldo(*gin.Context)
	// TransferHandler should handle balance mutation event
	Transfer(*gin.Context)
	TransferMinus(*gin.Context)
}

// New is a function for creating handler
func New(quorum int, controller c.Controller) Handler {
	return &handler{
		controller: controller,
		quorum: quorum,
	}
}

// Ping if a function for handling healthcheck in top level routing
func (h *handler) Ping(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	ctx.JSON(200, gin.H{
		"pingReturn": 1,
	})

	return
}

func (h * handler) Register(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	quorum := ctx.MustGet("quorum").(int)
	if quorum < (h.quorum/2) + 1 {
		ctx.JSON(200, gin.H{
			"transferReturn": -2,
		})
		return
	}

	var request m.RegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"registerReturn": -99,
		})
		return
	}

	customer := m.Customer{
		UserID: request.UserID,
		Name: request.Nama,
		Balance: 0,
	}

	_, err = h.controller.Register(ctx, customer)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"registerReturn": -4,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"registerReturn": 1,
	})

	return
}

func (h * handler) GetSaldo(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	quorum := ctx.MustGet("quorum").(int)
	if quorum < (h.quorum/2) + 1 {
		ctx.JSON(200, gin.H{
			"saldo": -2,
		})
		return
	}

	var request m.SaldoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"saldo": -99,
		})
		return
	}

	customer, err := h.controller.GetCustomer(ctx, request.UserID)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"saldo": -1,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"saldo": customer.Balance,
	})

	return
}

func (h * handler) GetTotalSaldo(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	quorum := ctx.MustGet("quorum").(int)
	if quorum < h.quorum {
		ctx.JSON(200, gin.H{
			"saldo": -2,
		})
		return
	}

	var request m.SaldoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"saldo": -99,
		})
		return
	}

	saldo, err := h.controller.GetTotalSaldo(ctx, request.UserID)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"saldo": saldo,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"saldo": saldo,
	})

	return
}

func (h * handler) Transfer(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	quorum := ctx.MustGet("quorum").(int)
	if quorum < (h.quorum/2) + 1 {
		ctx.JSON(200, gin.H{
			"transferReturn": -2,
		})
		return
	}

	var transfer m.TransferRequest
	err := ctx.ShouldBindJSON(&transfer)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"transferReturn": -99,
		})
		return
	}

	if transfer.Nilai > 10e9 {
		ctx.JSON(400, gin.H{
			"transferReturn": -5,
		})
		return
	}

	transferReturn, err := h.controller.Transfer(ctx, transfer.UserID, transfer.Nilai)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"transferReturn": transferReturn,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"transferReturn": transferReturn,
	})

	return
}

func (h * handler) TransferMinus(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	quorum := ctx.MustGet("quorum").(int)
	if quorum < (h.quorum/2) + 1 {
		ctx.JSON(200, gin.H{
			"transferReturn": -2,
		})
		return
	}

	var transfer m.TransferRequest
	err := ctx.ShouldBindJSON(&transfer)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{
			"transferReturn": -99,
		})
		return
	}

	if transfer.Nilai > 10e9 {
		ctx.JSON(400, gin.H{
			"transferReturn": -5,
		})
		return
	}

	transferReturn, err := h.controller.TransferMinus(ctx, transfer.UserID, transfer.Nilai)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"transferReturn": transferReturn,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"transferReturn": transferReturn,
	})

	return
}
