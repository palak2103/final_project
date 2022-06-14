package main

import (
	//"fmt"
	//"time"
	//"final_project/db"

	"github.com/gofiber/fiber/v2"

	_ "github.com/mattn/go-sqlite3"
)

//logout handler

func showLogoutHandler(c *fiber.Ctx) error {
	return c.Render("logout", fiber.Map{})
}

func LogoutHandler(c *fiber.Ctx) error {
	// cookie := fiber.Cookie{
	// 	Name:     "auth",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&cookie)

	// return c.JSON(fiber.Map{
	// 	"message": "you are logout",
	// })

		p := new(Logout)
	
		if err := c.BodyParser(p); err != nil {
			return err
		}
		
	
		return c.Render("logout", fiber.Map{
			"message" :"you are logout",
		})
	
}
	


