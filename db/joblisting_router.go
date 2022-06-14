package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//for job listing
// add path
func RouteJoblistings(app *fiber.App) {

	//Read all operation in json form on browser
	app.Get("/joblistings", func(c *fiber.Ctx) error {
		return c.JSON(GetJoblistings())

	})

	//Create  a new joblisting CREATE OPERATION
	app.Post("/joblistings", func(c *fiber.Ctx) error {
		joblisting := Joblisting{}
		if err := c.BodyParser(&joblisting); err != nil {
			return err
		}
		InsertJoblisting(joblisting)
		return c.JSON(GetJoblistings())
	})

	//Get a joblist by id READ OPERATION
	app.Get("/joblistings/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		joblisting, err := FindJoblisting(idInt)
		if err != nil {
			return err
		}
		return c.JSON(joblisting)
	})
	//Delete a job by id DELETE OPERATION

	app.Delete("/joblistings/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteJoblisting(idInt)
		return c.JSON(GetJoblistings())
	})

	// Update a joblisting by id UPDATE operation

	app.Put("/joblistings/:id", func(c *fiber.Ctx) error {
		joblisting := Joblisting{}

		if err := c.BodyParser(&joblisting); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		joblisting.ID = uint(idInt)
		UpdateJoblisting(joblisting)
		return c.JSON(GetJoblistings())

	})
}





