package main

import (
	"final_project/db"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/mattn/go-sqlite3"
)
// secretkey key var
var Secretkey = "secretkey"
// user var
var users = []User{

	{EmailId: "palakgupta120898@gmail.com", Name: "palak"},
}

func main() {
	db.Init()    //incilization
	db.Migrate() //Migration

	app := configApp()
	db.RouteJoblistings(app)
	db.RouteLoginPersons(app)
	db.RouteResisterPersons(app)
	db.RouteAdmins(app)
	db.RouteApplicants(app)
	db.RouteLogouts(app)
	app.Listen(":3000")  // listen localhost:3000
}
// configapp function
func configApp() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Get("/admin_dashboard",admin_dashboardHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})
	app.Post("/users", func(c *fiber.Ctx) error {
		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return err
		}
		users = append(users, *u)
		return c.JSON(users)
	})
	app.Get("/users/:emailid/:field?", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("emailid"))
		for _, user := range users {
			if user.EmailId == c.Params("emailid") {
				return c.JSON(user)
			}
		}
		return fiber.NewError(fiber.StatusNotFound, "this user not available:"+c.Params("password"))
	})

// ALL HANDLER MAPING HERE,

	//showlogin handler, login handler (map)
	app.Get("/login", showloginPersonHandler) // for get login method
	app.Post("/login", loginPersonHandler)    // for post login method

	// showRegistration handler , registration (map)

	app.Get("/registration", showRegisterPersonHandler) // for get register method
	app.Post("/registration", RegisterPersonHandler)    // for post register method

	// showuser handler , user(map)

	

	//showlogout handler logout (map)

	app.Get("/logout", showLogoutHandler)                 // for get logout method
	app.Post("/logout", LogoutHandler)                     // for post logout method

	// show applicant handler (map)
	app.Get("/applicant", showApplicantHandler)            // for get applicant  method
	app.Post("/applicant", ApplicantHandler)                 // for post applicant method

	// show admin handler (map)
	app.Get("/admin", showAdminHandler)                     // for get admin method
	app.Post("/admin", AdminHandler)                           // for post admin method
	
	// show joblisting handler(map)

	app.Get("/joblisting", showJoblistingHandler)           // for get joblisting method
	app.Post("/joblisting", JoblistingHandler)              // for post joblisting method
	app.Get("/linked_out", DashboardHandler)                 // for get linked_out page Dashboard 
	return app
}
// struct for regiterperson

type RegisterPerson struct {
	Name            string ` json:"name" xml:"name" form:"name" validate:"required"`
	EmailId         string `json:"emailid" gorm:"unique" xml:"emailid" form:"emailid"  validate:"required,emailid,min=6,max=32"`
	Password        string `json:"-" xml:"password" form:"password" validate:"required"`
	Confirmpassword string `json:"-" xml:"confirmpassword" form:"confirmpassword" validate:"required"`
}

// struct for loginperson with json return, xml form, and validation

type LoginPerson struct {
	EmailId  string `json:"emailid" gorm:"unique"  xml:"emailid" form:"emailid" validate:"required,emailid,min=6,max=32"`
	Password string `json:"-" xml:"password" form:"password" validate:"required"`
}
// struct for user with json return, xml form, and validation

type User struct {
	Id       uint   `json:"id" xml:"id" form:"id" validate:"required,number"`
	Name     string ` json:"name" xml:"name" form:"name" validate:"required"`
	EmailId  string `json:"emailid" gorm:"unique" xml:"emailid" form:"emailid"  validate:"required,emailid,min=6,max=32"`
	Password string `json:"-" xml:"password" form:"password" validate:"required"`
}

// struct for joblisting with json return, xml form, and validation

type Joblisting struct {
	Position          string `json:"position" xml:"position" form:"position" validation:"required"`
	Company           string `json:"company" xml:"company" form:"company" validation:"required"`
	Description       string `json:"description" xml:"description" form:"description" validation:"required"`
	Salaryexpectation int    `json:"salaryexpectation" xml:"salaryexpectation" form:"salaryexpectation" validation:"required number"`
}
// struct for applicant with json return, xml form, and validation

type Applicant struct {
	FirstName        string `json:"firstname"  xml:"firstname" form:"firstname" validation:"required"`
	LastName         string `json:"lastname"  xml:"lasttname" form:"lastname" validation:"required"`
	EmailId          string `json:"emailid"  xml:"emailid" form:"emailid" validation:"required"`
	CurrentAddress   string `json:"currentaddress"  xml:"currentaddress" form:"currentaddress" validation:"required"`
	PermanantAddress string `json:"permanantaddress"  xml:"permanantaddress" form:"permanantaddress" validation:"required"`
	Qualification    string `json:"qualification"  xml:"qualification" form:"qualification" validation:"required"`
	CurrentStatus    string `json:"currentstatus"  xml:"currentstatus" form:"currentstatus" validation:"required"`
	Company          string `json:"company"  xml:"company" form:"company" validation:"required"`
	Position         string `json:"position"  xml:"position" form:"position" validation:"required"`
	Skills            string `json:"skills"  xml:"skills" form:"skills" validation:"required"`



}

// struct for admin with json return, xml form, and validation

type Admin struct {
	EmailId    string `json:"emailid" xml:"emailid" form:"emailid" validation:"required"`
	Password   string `json:"password" xml:"password" form:"password" validation:"required"`
}



// struct for logout with json return, xml form, and validation

type Logout struct {
	EmailId string `json:"emailid" xml:"emailid" form:"emailid"`
	Password string `json:"password" xml:"password" form:"password"`
}

