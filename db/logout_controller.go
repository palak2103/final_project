package db
// patch all logout
func GetLogouts() []Logout {
	var logouts []Logout
	Db.Find(&logouts)
	return logouts
}

//BASIC OPERATIONS
// create operation (insert)

func InsertLogout(logout Logout) {
	Db.Create(&logout)
}

// Delete operation
func DeleteLogout(id int) {
	var logout Logout
	Db.First(&logout, id) //petch to id
	Db.Delete(&logout)    //delete
}

//Update data
func UpdateLogout(logout Logout) {
	llogout := Logout{} //to jjoblisting  we can mannually save data
	Db.First(&llogout, logout.ID)
	//fetch to id
	llogout.EmailId =logout.EmailId
	llogout.Password = logout.Password
	Db.Save(&logout)
}


// find logout user by id
func FindLogout(id int) (Logout, error) {
	var logout Logout
	err := Db.First(&logout, id).Error
	return logout, err
}