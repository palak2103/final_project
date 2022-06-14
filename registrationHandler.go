package main

import (
	//"log"
	"fmt"
	"final_project/db"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	 
)

// showregisterperson handler function
func showRegisterPersonHandler(c *fiber.Ctx) error {
	return c.Render("registration", fiber.Map{})
}
// registerperson handler functiom
func RegisterPersonHandler(c *fiber.Ctx) error {
	registerperson := new(RegisterPerson) //var

	err := c.BodyParser(registerperson)
	if err != nil {
		return err
	}

	user := User{
		Name:     registerperson.Name,
		EmailId:  registerperson.EmailId,
		Password: registerperson.Password,
	}
	// Create user condtition
	if err:= db.Db.Create(&user).Error  ; err != nil {
		fmt.Println("user not found" , err)

	}
	
  // return login page
	return c.RedirectToRoute("/login", fiber.Map{"message":"Registration succesfull"})
}





































