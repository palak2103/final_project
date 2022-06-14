package main

import (
	"fmt"
	//"log"
	//"github.com/go-playground/validator"
	//"strconv"
	//"time"

	//"log"
	"final_project/db"
	//"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"

	_ "github.com/mattn/go-sqlite3"
)

///LOGIN FUNCTION
//this function will show login page
func showloginPersonHandler(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func loginPersonHandler(c *fiber.Ctx) error {
	loginperson := new(LoginPerson)

	err := c.BodyParser(loginperson)

	if err != nil {
		return err
	}
	user := User{}

	err = db.Db.Where("email_id = ? and password = ?", loginperson.EmailId, loginperson.Password).First(&user).Error

	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    strconv.Itoa(int(user.Id)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// })

	// token, err := claims.SignedString([]byte(Secretkey))

	// if err != nil {
	// 	c.Status(fiber.StatusInternalServerError)
	// 	return c.JSON(fiber.Map{
	// 		"message": "could not login",
	// 	})
	// }

	// cookie := fiber.Cookie{
	// 	Name:     "auth",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&cookie)

	// return c.JSON(fiber.Map{
	// 	"message": "success",
	// })

	//print users
	//return user

	fmt.Println(user, "user login")

	//return c.JSON(loginperson)
	//return c.RedirectToRoute("linked_out", fiber.Map{"message": "Login succesfull"})
	return c.Redirect("/linked_out")
	//return c.Render("linked_out", fiber.Map{"message": "Login succesfull"})

}

