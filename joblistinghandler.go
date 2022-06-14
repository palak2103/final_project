package main

import (
	"fmt"

	"final_project/db"

	"github.com/gofiber/fiber/v2"

	_ "github.com/mattn/go-sqlite3"
)
// show joblisting handler
func showJoblistingHandler(c *fiber.Ctx) error {
	return c.Render("joblisting", fiber.Map{})
}
// joblisting handler
func JoblistingHandler(c *fiber.Ctx) error {
	joblisting := new(Joblisting)

	err := c.BodyParser(joblisting)
	if err != nil {
		return err
	}

	joblisting = &Joblisting{
		Position:          joblisting.Position,
		Company:           joblisting.Company,
		Description:       joblisting.Description,
		Salaryexpectation: joblisting.Salaryexpectation,
	}

	if err := db.Db.Create(&joblisting).Error; err != nil {
		fmt.Println("data not found", err)

	}

	//return linked_out page
	return c.Redirect("/linked_out")

}
//Dashboard handler
func DashboardHandler(c *fiber.Ctx) error {

	listings := db.GetJoblistings()
	// return posted jobs render in linked_out page
	return c.Render("linked_out", fiber.Map{"listings": listings})
}
