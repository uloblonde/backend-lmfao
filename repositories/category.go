package repositories

import (
	"backendtask/models"

	"gorm.io/gorm"
)

type CatRepository interface {
	CariCat() ([]models.Category, error)
	DapatCat(Id int) (models.Category, error)                       //Inisialisasi function cari user untuk mendapatkan models dari user bedasarkan ID
	MembuatCat(cat models.Category) (models.Category, error)        //inisialisasi membuat Category
	UpdateCat(cat models.Category, Id int) (models.Category, error) //Update Category
	HapusCat(cat models.Category, Id int) (models.Category, error)  //Hapuse User
}
type catRepo struct {
	db *gorm.DB //ngepointer si gorm db
}

func RepositoryCat(db *gorm.DB) *repo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &repo{db}
}
func (r *repo) CariCat() ([]models.Category, error) {
	var category []models.Category
	err := r.db.Find(&category).Error // Using Find method
	return category, err
}
func (r *repo) DapatCat(Id int) (models.Category, error) {
	var cat models.Category
	err := r.db.First(&cat, Id).Error
	return cat, err
}

func (r *repo) MembuatCat(cat models.Category) (models.Category, error) {
	err := r.db.Create(&cat).Error
	return cat, err
}

func (r *repo) UpdateCat(cat models.Category, Id int) (models.Category, error) {
	err := r.db.Save(&cat).Error
	return cat, err
}

func (r *repo) HapusCat(cat models.Category, Id int) (models.Category, error) {
	err := r.db.Delete(&cat).Error
	return cat, err
}
