package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"lifedev/shop/controllers"
)

func main() {
	// session
	// store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")

	// controllers
	// helloController := controllers.InitHelloController(store)
	prodController := controllers.InitProductController()
	tranController := controllers.InitTransactionController()
	// authController := controllers.InitAuthController(store)

	prod := app.Group("/products")
	prod.Get("/", prodController.HomeProduct)
	prod.Get("/dashboard", prodController.DashboardProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", prodController.EditlProduct)
	prod.Post("/editproduct/:id", prodController.EditlPostedProduct)
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	tran := app.Group("/transactions")
	tran.Get("/", tranController.DashboardTransaction)
	tran.Post("/create", tranController.AddPostedTransaction)
	tran.Get("/delete/:id", tranController.DeleteTransactionById)

	// app.Get("/login", authController.Login)
	// app.Post("/login", authController.LoginPosted)
	// app.Get("/logout", authController.Logout)
	//app.Get("/profile",authController.Profile)

	// app.Use("/profile", func(c *fiber.Ctx) error {
	// 	sess,_ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// })
	// app.Get("/profile", func(c *fiber.Ctx) error {
	// 	sess,_ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// }, authController.Profile)

	app.Listen(":3000")
}
