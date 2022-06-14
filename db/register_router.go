package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)
//ROUTING FOR REGISTRATION
func RouteResisterPersons(app *fiber.App) {

	//Read all operation for registerpersons in json form on browser
	app.Get("/;registerpersons", func(c *fiber.Ctx) error {
		return c.JSON(GetRegisterPersons())
	})
	//Insert operation for registerperson

	app.Post("/registerpersons", func(c *fiber.Ctx) error {
		registerperson := RegisterPerson{}
		if err := c.BodyParser(&registerperson); err != nil {
			return err
		}
		InsertRegisterPerson(registerperson)
		return c.JSON(GetRegisterPersons())
	})

	//READ  a regiterperson by id
	app.Get("/registerpersons/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		registerperson, err := FindRegisterPerson(idInt)
		if err != nil {
			return err
		}
		return c.JSON(registerperson)
	})

	//Delete operation for registerperson  by id
	app.Delete("/registerpersons/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteRegisterPerson(idInt)
		return c.JSON(GetRegisterPersons())
	})

	//Update operation for register person by id
	app.Put("/registerpersons/:id", func(c *fiber.Ctx) error {
		registerperson := RegisterPerson{}

		if err := c.BodyParser(&registerperson); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		registerperson.ID = uint(idInt)
		UpdateRegisterPerson(registerperson)
		return c.JSON(GetRegisterPersons())

	})
}