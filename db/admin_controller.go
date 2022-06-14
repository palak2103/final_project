package db
// CRUD operations for admin (admin_service)
// For a get 
func GetAdmins() []Admin {
	var admins []Admin
	Db.Find(&admins)
	return admins
}
// Create operation
func InsertAdmin(admin Admin) {
	Db.Create(&admin)
}
// Delete operation
func DeleteAdmin(id int) {
	var admin Admin
	Db.First(&admin, id)
	Db.Delete(&admin)
}
// update operation
func UpdateAdmin(admin Admin) {
	aadmin := Admin{}
	Db.First(&aadmin, admin.ID)
	aadmin.EmailId = admin.EmailId
	aadmin.Password = admin.Password
	Db.Save(&admin)
}
// find operation
func FindAdmin(id int) (Admin, error) {
	var admin Admin
	err := Db.First(&admin, id).Error
	return admin, err
}
