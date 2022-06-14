package db 
 //Get applicants operation
func GetApplicants() []Applicant {
	var applicants []Applicant
	Db.Find(&applicants)
	return applicants
}

//BASIC OPERATIONS
// create operation (insert) for applicant
func InsertApplicant(applicant Applicant) {
	Db.Create(&applicant)
}

// delete operation for applicant
func DeleteApplicant(id int) {
	var applicant Applicant
	Db.First(&applicant, id)
	Db.Delete(&applicant)
}

//update operation for applicant
func UpdateApplicant(applicant Applicant) {
	aapplicant := Applicant{}
	Db.First(&aapplicant, applicant.ID)
	aapplicant.FirstName = applicant.FirstName
	aapplicant.LastName = applicant.LastName
	aapplicant.EmailId = applicant.EmailId
	aapplicant.CurrentAddress = applicant.CurrentAddress
	aapplicant.PermanantAddress = applicant.PermanantAddress
	aapplicant.Qualification = applicant.Qualification
	aapplicant.CurrentStatus = applicant.CurrentStatus
	aapplicant.Company=applicant.Company
	aapplicant.Position =applicant.Position
	aapplicant.Skills = applicant.Skills
	Db.Save(&applicant)

}

// find operation for applicant
func FindApplicant(id int) (Applicant, error) {
	var applicant Applicant
	err := Db.First(&applicant, id).Error
	return applicant, err
}