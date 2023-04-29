package models

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email" binding:"required,name" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	FullName  string `json:"fullName" gorm:"type:varchar(255)"`
	Gender    string `json:"gender" gorm:"type:varchar(255)"`
	Phone     string `json:"phone" gorm:"type:varchar(255)"`
	Address   string `json:"address" gorm:"type:varchar(255)"`
	Role      string `json:"role" gorm:"type:varchar(255)"`
	Subscribe bool   `json:"subscribe" gorm:"default: false"`
}
type UserProfileResponse struct {
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	FullName string `json:"fullName" gorm:"type:varchar(255)"`
	Gender   string `json:"gender" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255)"`
	Address  string `json:"address" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(255)"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
