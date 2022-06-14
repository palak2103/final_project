package db

import (
	"fmt"
	//"strconv"

	//"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB //variable

// it will manage gorm (table)
type Joblisting struct {
	gorm.Model // composition design pattern

	Position          string `json:"position"`
	Company           string `json:"company"`
	Description       string `json:"description"`
	Salaryexpectation int    `json:"salaryexpectation"`
}

//
type Applicant struct {
	gorm.Model // composition design pattern

	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	EmailId          string `json:"emailid"`
	ContactNo        string `json:"contactno"`
	CurrentAddress   string `json:"currentaddress"`
	PermanantAddress string `json:"permanantaddress"`
	Qualification    string `json:"qualification"`
	CurrentStatus    string `json:"currentstatus"`
	Company          string `json:"company"`
	Position         string  `json:"position"`
	Skills            string `json:"skills"`

}
// struct for regiterperson
type RegisterPerson struct {
	gorm.Model
	Name     string `json:"name"`
	EmailId  string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

// struct for loginperson

type LoginPerson struct {
	gorm.Model
	EmailId  string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
// struct for admin

type Admin struct {
	gorm.Model
	EmailId  string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}




// struct for logout

type Logout struct {
	gorm.Model // composition design pattern

	EmailId           string `json:"emailid"`
	Password           string `json:"password"` 

}

// For data base connection stablish
// data initilization
func Init() {
	db, err := gorm.Open(sqlite.Open("joblistings.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = db
	fmt.Println("DB Initilized")
}

// for create the table
//auto-change
func Migrate() {
	err := Db.AutoMigrate(&Joblisting{}, Applicant{}, LoginPerson{}, RegisterPerson{}, Admin{}, Logout{}) // table
	if err != nil {
		panic(err)
	}
	fmt.Println("DB migrated")

}
