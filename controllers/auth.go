package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"lifedev/shop/database"
	"lifedev/shop/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type LoginController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}
type AuthController struct {
	// declare variables
	store *session.Store
}

func InitAuthController(s *session.Store) *LoginController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})
	return &LoginController{Db: db, store: s}
}

// GET /login
func (controller *LoginController) Login(c *fiber.Ctx) error {
	// load all products

	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// GET /register
func (controller *LoginController) Register(c *fiber.Ctx) error {
	// load all products

	return c.Render("register", fiber.Map{
		"Title": "Registerasi",
	})
}

// POST /register
func (controller *LoginController) AddPostedRegister(c *fiber.Ctx) error {
	// load all products
	var myform models.User

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	// save product
	err := models.CreateUser(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/login")
	}
	// if succeed
	return c.Redirect("/login")
}

// POST /login
func (controller *LoginController) LoginPosted(c *fiber.Ctx) error {
	// load all products
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	var product models.Product
	err := models.ReadUser(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	var myform models.User

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	// hardcode auth
	if myform.Username == "admin" && myform.Password == "1234" {
		sess.Set("username", "admin")
		sess.Save()

		return c.Redirect("/products")
	}
	return c.Redirect("/login")

}

// /profile

func (controller *LoginController) Profile(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.JSON(fiber.Map{
		"username": val,
	})
}

// /logout
func (controller *LoginController) Logout(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()

	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
