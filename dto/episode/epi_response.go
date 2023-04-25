package episdto

import "backendtask/models"

type EpisodeResponse struct {
	Id            int         `json:"id"`
	Title         string      `json:"title" gorm:"type:varchar(255)"`
	ThumbnailFilm string      `json:"thumbnailFilm" gorm:"type:varchar(255)"`
	LinkFilm      string      `json:"linkFilm" gorm:"type:varchar(255)"`
	Film          models.Film `json:"film"`
	FilmId        int         `json:"filmId"`
}
