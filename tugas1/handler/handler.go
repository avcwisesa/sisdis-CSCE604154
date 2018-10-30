package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"

	m "github.com/avcwisesa/sisdis/tugas1/model"
)

// handler holds the structure for Handler
type handler struct {}

// Handler holds the contract for Handler
type Handler interface {
	// Ping should handle healthcheck in top level routing
	Ping(*gin.Context)
	// RegisterHandler should handle new customer register
	RegisterHandler(*gin.Context)
	// GetSaldoHandler should handle get balance
	GetSaldoHandler(*gin.Context)
	// GetTotalSaldoHandler should handle get balance from all branches
	GetTotalSaldoHandler(*gin.Context)
	// TransferHandler should handle balance mutation event
	TransferHandler(*gin.Context)
}

// New is a function for creating handler
func New() Handler {
	return &handler{}
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

	ctx.JSON(200, gin.H{
		"pingReturn": 1,
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

	var saldo m.SaldoRequest
	err := ctx.ShouldBindJSON(&saldo)
	if err != nil {
		ctx.JSON(400, gin.H{
			"saldo": -99,
		})
		return
	}

	userID, err := strconv.Atoi(saldo.UserID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"saldo": -99,
		})
		return
	}

	saldo, err := c.GetSaldo(userID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"saldo": -4,
		})
	}

	ctx.JSON(200, gin.H{
		"saldo": saldo,
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

	var saldo m.SaldoRequest


	// check domisili

	ctx.JSON(200, gin.H{
		"saldo": 1,
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

	var transfer m.TransferRequest
	err := ctx.ShouldBindJSON(&transfer)
	iff err != nil {
		ctx.JSON(400, gin.H{
			"transferReturn": -99,
		})
		return
	}

	// check transfer value, if x < 0 || > 1M --> -5

	// get user from db, else --> -1

	// db error --> -4

	ctx.JSON(200, gin.H{
		"transferReturn": 1,
	})

	return
}