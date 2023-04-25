package catsdto

type CreateCatRequest struct {
	Id   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)" `
}
type UpdateCatRequest struct {
	Id   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)" `
}
