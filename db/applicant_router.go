package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)
// Routing for applicants
func RouteApplicants(app *fiber.App) {
	app.Get("/apply", func(c *fiber.Ctx) error {
		return c.Render("applicant", fiber.Map{})

	})


	//Read all operation in json form on browser
	app.Get("/applicants", func(c *fiber.Ctx) error {
		return c.JSON(GetApplicants())

	})
	
    // Insert applicants operation
	app.Post("/applicants", func(c *fiber.Ctx) error {
		applicant := Applicant{}
		if err := c.BodyParser(&applicant); err != nil {
			return err
		}
		InsertApplicant(applicant)
		return c.JSON(GetApplicants())
	})
    //Get aaplicant
	app.Get("/applicants/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		applicant, err := FindApplicant(idInt)
		if err != nil {
			return err
		}
		return c.JSON(applicant)
	})
	// Delete applicant operation

	app.Delete("/applicants/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		DeleteApplicant(idInt)
		return c.JSON(GetApplicants())
	})

	//update applicants operation

	app.Put("/applicants/:id", func(c *fiber.Ctx) error {
		applicant := Applicant{}

		if err := c.BodyParser(&applicant); err != nil {
			return err
		}
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		applicant.ID = uint(idInt)
		UpdateApplicant(applicant)
		return c.JSON(GetApplicants())

	})
}





