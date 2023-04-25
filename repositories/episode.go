package repositories

import (
	"backendtask/models"
	"fmt"

	"gorm.io/gorm"
)

type EpiRepository interface {
	CariEpi() ([]models.Episode, error)
	DapatEpi(Id int) (models.Episode, error)                      //Inisialisasi function cari user untuk mendapatkan models dari user bedasarkan ID
	MembuatEpi(cat models.Episode) (models.Episode, error)        //inisialisasi membuat Episode
	UpdateEpi(cat models.Episode, Id int) (models.Episode, error) //Update Episode
	HapusEpi(cat models.Episode, Id int) (models.Episode, error)  //Hapuse User
	CariEpiByFilm(Id int) ([]models.Episode, error)
	DapatEpiByFilm(Id int, IDE int) (models.Episode, error)

	// DapatCatId(Id int) (models.Category, error)
}
type epiRepo struct {
	db *gorm.DB //ngepointer si gorm db
}

func RepositoryEpi(db *gorm.DB) *epiRepo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &epiRepo{db}
}
func (r *epiRepo) CariEpiByFilm(Id int) ([]models.Episode, error) {
	var Episode []models.Episode
	err := r.db.Preload("Film.Category").Find(&Episode, "film_id = ?", Id).Error

	fmt.Println(Id)
	return Episode, err
}

func (r *epiRepo) DapatEpiByFilm(Id int, IDE int) (models.Episode, error) {
	var Episode models.Episode
	err := r.db.Preload("Film").First(&Episode, "film_id = ? AND id = ?", Id, IDE).Error

	return Episode, err
}

func (r *epiRepo) CariEpi() ([]models.Episode, error) {
	var episode []models.Episode
	err := r.db.Preload("Film").Find(&episode).Error // Using Find method
	return episode, err
}
func (r *epiRepo) DapatEpi(Id int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film").First(&episode, Id).Error
	return episode, err
}

func (r *epiRepo) MembuatEpi(epi models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Create(&epi).Error
	return epi, err
}

func (r *epiRepo) UpdateEpi(epi models.Episode, Id int) (models.Episode, error) {
	err := r.db.Save(&epi).Error
	return epi, err
}

func (r *epiRepo) HapusEpi(epi models.Episode, Id int) (models.Episode, error) {
	err := r.db.Delete(&epi).Error
	return epi, err
}

// func (r *repo) GetCatId(Id int) (models.Category, error) {
// 	var cat models.Category
// 	err := r.db.First(&cat, Id).Error
// 	return cat, err
// }
