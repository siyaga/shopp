package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"lifedev/shop/database"
	"lifedev/shop/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type TransactionController struct {
	// declare variables
	Db *gorm.DB
}

func InitTransactionController() *TransactionController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Transaction{})

	return &TransactionController{Db: db}
}

// routing
// GET /transactions
func (controller *TransactionController) DashboardTransaction(c *fiber.Ctx) error {
	// load all products
	var transactions []models.Transaction
	err := models.ReadTransaction(controller.Db, &transactions)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("transaction", fiber.Map{
		"Title":        "Daftar Produk",
		"Transactions": transactions,
	})
}

// POST /products/create
func (controller *TransactionController) AddPostedTransaction(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Transaction

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/transactions")
	}
	// save product
	err := models.CreateTransaction(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/transactions")
	}
	// if succeed
	return c.Redirect("/transactions")
}

// / GET /products/deleteproduct/xx
func (controller *TransactionController) DeleteTransactionById(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var transactions models.Transaction
	models.DeleteTransactionById(controller.Db, &transactions, idn)
	return c.Redirect("/transactions")
}
