package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Data patching
// Router for admin
func RouteAdmins(app *fiber.App) {

	//Read all operation in json form browser
	app.Get("/;admins", func(c *fiber.Ctx) error {
		return c.JSON(GetAdmins())
	})
		// Create a new CREATE operation


	app.Post("/admins", func(c *fiber.Ctx) error {
		admin := Admin{}
		if err := c.BodyParser(&admin); err != nil {
			return err
		}
		InsertAdmin(admin)
		return c.JSON(GetAdmins())
	})
	app.Get("/admins/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		admin, err := FindAdmin(idInt)
		if err != nil {
			return err
		}
		return c.JSON(admin)
	})
	app.Delete("/admins/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteAdmin(idInt)
		return c.JSON(GetAdmins())
	})
	app.Put("/admins/:id", func(c *fiber.Ctx) error {
		admin := Admin{}

		if err := c.BodyParser(&admin); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		admin.ID = uint(idInt)
		UpdateAdmin(admin)
		return c.JSON(GetAdmins())

	})
	
}