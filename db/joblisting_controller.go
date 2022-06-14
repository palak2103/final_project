package db

// patch all
func GetJoblistings() []Joblisting {
	var joblistings []Joblisting
	Db.Find(&joblistings)
	return joblistings
}

//BASIC OPERATIONS
// create operation (insert)

func InsertJoblisting(joblisting Joblisting) {
	Db.Create(&joblisting)
}

// Delete operation
func DeleteJoblisting(id int) {
	var joblisting Joblisting
	Db.First(&joblisting, id) //petch to id
	Db.Delete(&joblisting)    //delete
}

//Update joblisting operation
func UpdateJoblisting(joblisting Joblisting) {
	jjoblisting := Joblisting{}            //to jjoblisting  we can mannually save data
	Db.First(&jjoblisting, joblisting.ID)  //fetch to id
	jjoblisting.Position = joblisting.Position
	jjoblisting.Company = joblisting.Company
	jjoblisting.Description = joblisting.Description
	jjoblisting.Salaryexpectation = joblisting.Salaryexpectation
	Db.Save(&joblisting)
}

//find a joblisting

func FindJoblisting(id int) (Joblisting, error) {
	var joblisting Joblisting
	err := Db.First(&joblisting, id).Error
	return joblisting, err
}









