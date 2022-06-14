package main

import (
	"fmt"

	"final_project/db"

	"github.com/gofiber/fiber/v2"

	_ "github.com/mattn/go-sqlite3"
)

func showApplicantHandler(c *fiber.Ctx) error {
	return c.Render("applicant", fiber.Map{})
}
func ApplicantHandler(c *fiber.Ctx) error {
	applicant := new(Applicant)

	err := c.BodyParser(applicant)
	if err != nil {
		return err
	}
	applicant = &Applicant{
		FirstName:        applicant.FirstName,
		LastName:         applicant.LastName,
		CurrentAddress:   applicant.CurrentAddress,
		PermanantAddress: applicant.PermanantAddress,
		EmailId:          applicant.EmailId,
		CurrentStatus:    applicant.CurrentStatus,
		Qualification:    applicant.Qualification,
		Position:         applicant.Position,
		Company:          applicant.Company,
		Skills:           applicant.Skills,
	}
	if err := db.Db.Create(&applicant).Error; err != nil {
		fmt.Println("applicant not found")

	}
	//return c.Redirect("/applicant")
	return c.Redirect("/linked_out")


}
