package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)
// Routing for login persons
func RouteLoginPersons(app *fiber.App) {

	// Read operation all login person in json form on browser

	app.Get("/;loginpersons", func(c *fiber.Ctx) error {
		return c.JSON(GetLoginPersons())
	})
	// Insert operation for login person

	app.Post("/loginpersons", func(c *fiber.Ctx) error {
		loginperson := LoginPerson{}
		if err := c.BodyParser(&loginperson); err != nil {
			return err
		}
		InsertLoginPerson(loginperson)
		return c.JSON(GetLoginPersons())
	})

	//Read operation for login person by id
	app.Get("/loginpersons/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		loginperson, err := FindLoginPerson(idInt)
		if err != nil {
			return err
		}
		return c.JSON(loginperson)
	})
	//Delete operation for login person by id
	app.Delete("/loginpersons/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteLoginPerson(idInt)
		return c.JSON(GetLoginPersons())
	})
	
	//Update operation for login person by id
	app.Put("/loginpersons/:id", func(c *fiber.Ctx) error {
		loginperson := LoginPerson{}

		if err := c.BodyParser(&loginperson); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		loginperson.ID = uint(idInt)
		UpdateLoginPerson(loginperson)
		return c.JSON(GetLoginPersons())

	})
}