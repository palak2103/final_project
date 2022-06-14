package db
// patch all loginperson
// add path for login persons

//Read all login persons
func GetLoginPersons() []LoginPerson {
	var loginpersons []LoginPerson
	Db.Find(&loginpersons)
	return loginpersons
}
//Insert operation for login person 
func InsertLoginPerson(loginperson LoginPerson) {
	Db.Create(&loginperson)
}
// Delete operation for login person by id

func DeleteLoginPerson(id int) {
	var loginperson LoginPerson
	Db.First(&loginperson, id)
	Db.Delete(&loginperson)
}
// update operation for login person 
func UpdateLoginPerson(loginperson LoginPerson) {
	lloginperson := LoginPerson{}
	Db.First(&lloginperson, loginperson.ID)
	lloginperson.EmailId = loginperson.EmailId
	lloginperson.Password = loginperson.Password
	Db.Save(&loginperson)
}
// find operation for login person by id 

func FindLoginPerson(id int) (LoginPerson, error) {
	var loginperson LoginPerson
	err := Db.First(&loginperson, id).Error
	return loginperson, err
}