package repositories

import (
	"backendtask/models"

	"gorm.io/gorm"
)

type TransRepository interface {
	CariTrans() ([]models.Transaction, error)
	MembuatTrans(Trans models.Transaction) (models.Transaction, error)
	DapatTrans(Id int) (models.Transaction, error)
	UpdateTrans(status string, orderId int) (models.Transaction, error)
	HapusTrans(Trans models.Transaction, Id int) (models.Transaction, error)
}
type transRepo struct {
	db *gorm.DB //ngepointer si gorm db
}

func RepositoryTrans(db *gorm.DB) *transRepo { //function Repository mengambil parameter berupa pointer ke gorm dan mengembalikan pointer ke repo
	return &transRepo{db}
}
func (r *transRepo) CariTrans() ([]models.Transaction, error) {
	var trans []models.Transaction
	err := r.db.Preload("User").Find(&trans).Error // Using Find method
	return trans, err
}
func (r *transRepo) DapatTrans(Id int) (models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Preload("User").First(&trans, Id).Error
	return trans, err
}
func (r *transRepo) MembuatTrans(trans models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&trans).Error
	return trans, err
}
func (r *transRepo) UpdateTrans(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.First(&user, transaction.UserId)
		user.Subscribe = true
		r.db.Save(&user)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *transRepo) HapusTrans(trans models.Transaction, Id int) (models.Transaction, error) {
	err := r.db.Delete(&trans).Error
	return trans, err
}
