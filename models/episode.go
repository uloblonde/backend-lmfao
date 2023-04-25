package models

type Episode struct {
	Id            int    `json:"id"`
	Title         string `json:"title" gorm:"type:varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" gorm:"type:varchar(255)"`
	LinkFilm      string `json:"linkFilm" gorm:"type:varchar(255)"`
	Film          Film   `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FilmId        int    `json:"filmId"`
}
