package models

type Film struct {
	Id            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" gorm:"type: varchar(255)" `
	ThumbnailFilm string   `json:"thumbnailFilm" gorm:"type: varchar(255)" `
	Year          int      `json:"year" `
	Category      Category `json:"category" `
	CategoryId    int      `json:"categoryId"`
	Description   string   `json:"description" gorm:"type: varchar(255)" `
}
type FilmResponse struct {
	Id            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" gorm:"type: varchar(255)" `
	ThumbnailFilm string   `json:"thumbnailFilm" gorm:"type: varchar(255)" `
	Year          int      `json:"year" `
	Category      Category `json:"category" `
	CategoryId    int      `json:"categoryId"`
	Description   string   `json:"description" gorm:"type: varchar(255)" `
}

func (FilmResponse) TableName() string {
	return "film_response"
}
