package controllers

import (
	"fmt"
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

type ProductController struct {
	// declare variables
	Db *gorm.DB
}

func InitProductController() *ProductController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}

// routing
// GET /products
func (controller *ProductController) HomeProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("home", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
	})
}

// routing
// GET /products
func (controller *ProductController) DashboardProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("products", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
	})
}

// GET /products/create
func (controller *ProductController) AddProduct(c *fiber.Ctx) error {
	return c.Render("addproduct", fiber.Map{
		"Title": "Tambah Produk",
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Product

	file, errFile := c.FormFile("image")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
		if errSaveFile != nil {
			fmt.Println("Fail to store file into public/ikmages directory.")
		}
	} else {
		fmt.Println("Nothing file to uploading.")
	}

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}

	myform.Image = filename
	// save product
	errr := models.CreateProduct(controller.Db, &myform)
	if errr != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/products")
}

// GET /products/productdetail?id=xxx
func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// GET /products/detail/xxx
func (controller *ProductController) GetDetailProduct2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// / GET products/editproduct/xx
func (controller *ProductController) EditlProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editproduct", fiber.Map{
		"Title":   "Edit Produk",
		"Product": product,
	})
}

// / POST products/editproduct/xx
func (controller *ProductController) EditlPostedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
		if errSaveFile != nil {
			fmt.Println("Fail to store file into public/ikmages directory.")
		}
	} else {
		fmt.Println("Nothing file to uploading.")
	}
	myform.Image = filename
	product.Name = myform.Name
	product.Image = myform.Image
	product.Description = myform.Description
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	product.Status = myform.Status
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.Redirect("/products")

}

// / GET /products/deleteproduct/xx
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	models.DeleteProductById(controller.Db, &product, idn)
	return c.Redirect("/products")
}
