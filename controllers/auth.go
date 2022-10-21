package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
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

type LoginForm struct {
	// declare variables
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
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
	var convertpass LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	comvertpassword, _ := bcrypt.GenerateFromPassword([]byte(convertpass.Password), 10)
	sHash := string(comvertpassword)

	myform.Password = sHash

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

	var user models.User
	var myform LoginForm
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	er := models.FindByUsername(controller.Db, &user, myform.Username)
	if er != nil {
		return c.Redirect("/login") // http 500 internal server error
	}

	// hardcode auth
	mycompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if mycompare != nil {
		sess.Set("username", user.Username)
		sess.Set("userID", user.Id)
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
