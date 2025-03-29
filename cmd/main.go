// package main

// import (
// 	"go-docker-code/docker"
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	docker.InitDockerClient()

// 	app := fiber.New()

// 	// API marshrutlari
// 	app.Get("/containers", ListContainers)
// 	app.Post("/containers", CreateContainer)
// 	app.Post("/containers/:id/start", StartContainer)
// 	app.Post("/containers/:id/stop", StopContainer)
// 	app.Delete("/containers/:id", RemoveContainer)

//		app.Listen(":8080")
//	}
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("hello Asadbek")
}
func Parameters(c *fiber.Ctx) error {
	return c.SendString("value" + c.Params("value"))
}
func API(c *fiber.Ctx) error {
	if c.Params("*") == "users" {
		return c.SendString("Hello, User endpoint")

	}
	if c.Params("*") == "file" {
		return c.SendString("Hello file?")
	}
	return c.SendString("404")
}
func main() {
	app := fiber.New()
	v1 := app.Group("api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, version v1")
	})
	v2 := app.Group("api/v2")
	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, version v2")
	})

	log.Fatal(app.Listen(":3000"))
	fmt.Println("Strated server")
}
