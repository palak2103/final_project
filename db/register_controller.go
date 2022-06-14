package db
// patch all Registerperson
// add path for register persons

//Read all registerperson
func GetRegisterPersons() []RegisterPerson {
	var registerpersons []RegisterPerson
	Db.Find(&registerpersons)
	return registerpersons
}
// Insert operation for registerperson 
func InsertRegisterPerson(registerperson RegisterPerson) {
	Db.Create(&registerperson)
}
//Delete operation for registerperson by id
func DeleteRegisterPerson(id int) {
	var registerperson RegisterPerson
	Db.First(&registerperson, id)
	Db.Delete(&registerperson)
}
//Update operation for registerperson 
func UpdateRegisterPerson(registerperson RegisterPerson) {
	rrgisterperson := RegisterPerson{}
	Db.First(&rrgisterperson, registerperson.ID)
	rrgisterperson.EmailId = registerperson.EmailId
	rrgisterperson.Password = registerperson.Password
	rrgisterperson.Name = registerperson.Name

	Db.Save(&registerperson)
}
//Find operation for registerperson by id
func FindRegisterPerson(id int) (RegisterPerson, error) {
	var registerperson RegisterPerson
	err := Db.First(&registerperson, id).Error
	return registerperson, err
}
