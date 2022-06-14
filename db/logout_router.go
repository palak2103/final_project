package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)
// Routing for logouts
func RouteLogouts(app *fiber.App) {
	app.Get("/logouts", func(c *fiber.Ctx) error {
		return c.JSON(GetLogouts())
	})
	//insert operation for logout in json form on browser
	app.Post("/logout", func(c *fiber.Ctx) error {
		logout := Logout{}
		if err := c.BodyParser(&logout); err != nil {
			return err
		}
		InsertLogout(logout)
		return c.JSON(GetLogouts())
	})
	//Read operation for logouts by id

	app.Get("/logouts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		logout, err := FindLogout(idInt)
		if err != nil {
			return err
		}
		return c.JSON(logout)
	})

	// DELETE operation by id

	app.Delete("/logouts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteLogout(idInt)
		return c.JSON(GetLogouts())
	})
	//put operation for logout by id

	app.Put("/logouts/:id", func(c *fiber.Ctx) error {
		logout := Logout{}

		if err := c.BodyParser(&logout); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		logout.ID = uint(idInt)
		UpdateLogout(logout)
		return c.JSON(GetLogouts())

	})

}
