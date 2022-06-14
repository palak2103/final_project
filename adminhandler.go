package main
import(
//	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
// //this function will show login page

func admin_dashboardHandler(c *fiber.Ctx) error {
	return c.Render("admin_dashboard", fiber.Map{
		"message": "welcome you are in admin dashboard",
		"users":   users,
	})
}
func showAdminHandler(c *fiber.Ctx) error {
	return c.Render("admin", fiber.Map{})
}

func AdminHandler(c *fiber.Ctx) error {
	p := new(Admin)

	if err := c.BodyParser(p); err != nil {
		return err
	}
	log.Println(p.EmailId)
	log.Println(p.Password)

	if emailid := p.EmailId; emailid == "palak120898@gmail.com" && p.Password == "palak"  {
		return c.Render("admin_dashboard", fiber.Map{
			"message": "login success",

		})
	}

	return c.Render("admin_dashboard", fiber.Map{
		"message" :"admin_dashboard",
	})

}

	



