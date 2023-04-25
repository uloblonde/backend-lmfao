package episdto

type CreateEpisodeRequest struct {
	Title         string `json:"title" form:"title" gorm:"type:varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type:varchar(255)"`
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type:varchar(255)"`
	FilmId        int    `json:"filmId" form:"filmId"`
}
type UpdateEpisodeRequest struct {
	Title         string `json:"title" form:"title" gorm:"type:varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type:varchar(255)"`
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type:varchar(255)"`
	FilmId        int    `json:"filmId" form:"filmId"`
}
