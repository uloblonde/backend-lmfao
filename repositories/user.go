package repositories

import (
	"backendtask/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CariUser() ([]models.User, error)                         //Inisialisasi function cari user untuk mendapatkan models dari user
	CariProfile(userId int) (models.User, error)              //Inisialisasi function cari user untuk mendapatkan models dari user
	DapatUser(Id int) (models.User, error)                    //Inisialisasi function cari user untuk mendapatkan models dari user bedasarkan ID
	MembuatUser(user models.User) (models.User, error)        //inisialisasi membuat user
	UpdateUser(user models.User, Id int) (models.User, error) //Update User
	HapusUser(user models.User, Id int) (models.User, error)  //Hapuse User

}
type repo struct {
	db *gorm.DB //ngepointer si gorm db
}

func RepositoryUser(db *gorm.DB) *repo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &repo{db}
}
func (r *repo) CariUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *repo) CariProfile(Id int) (models.User, error) {
	var profile models.User
	err := r.db.First(&profile, "id=?", Id).Error

	return profile, err
}

func (r *repo) DapatUser(Id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, Id).Error
	return user, err
}

func (r *repo) MembuatUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repo) UpdateUser(user models.User, Id int) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repo) HapusUser(user models.User, Id int) (models.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}

// func (r *repo) CariUser() ([]models.User, error) {
// 	var users []models.User
// 	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
// 	return users, err
// }
