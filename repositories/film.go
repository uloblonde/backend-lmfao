package repositories

import (
	"backendtask/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	CariFilm() ([]models.Film, error)
	MembuatFilm(film models.Film) (models.Film, error)
	DapatFilm(Id int) (models.Film, error)
	UpdateFilm(film models.Film, Id int) (models.Film, error) //Update Film
	HapusFilm(film models.Film, Id int) (models.Film, error)
	DapatCatId(Id int) (models.Category, error)
}
type filmRepo struct {
	db *gorm.DB //ngepointer si gorm db
}

func RepositoryFilm(db *gorm.DB) *repo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &repo{db}
}
func (r *repo) CariFilm() ([]models.Film, error) {
	var film []models.Film
	err := r.db.Preload("Category").Find(&film).Error // Using Find method
	return film, err
}
func (r *repo) DapatFilm(Id int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, Id).Error
	return film, err
}
func (r *repo) MembuatFilm(film models.Film) (models.Film, error) {
	err := r.db.Preload("Category").Create(&film).Error
	return film, err
}
func (r *repo) UpdateFilm(film models.Film, Id int) (models.Film, error) {
	err := r.db.Model(&film).Updates(&film).Error
	return film, err
}

func (r *repo) HapusFilm(film models.Film, Id int) (models.Film, error) {
	err := r.db.Delete(&film).Error
	return film, err
}
func (r *repo) DapatCatId(Id int) (models.Category, error) {
	var cat models.Category
	err := r.db.First(&cat, Id).Error
	return cat, err
}
