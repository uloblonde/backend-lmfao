package filmsdto

type CreatedFilmRequest struct {
	Id            int    `json:"id"`
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailFilm" validate:"required"`
	Year          int    `json:"year" gorm:"type:int(4)" form:"year" validate:"required"`
	CategoryId    int    `json:"categoryId" form:"categoryId" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
}
type UpdateFilmRequest struct {
	Id            int    `json:"id"`
	Title         string `json:"title" `
	ThumbnailFilm string `json:"thumbnailFilm" `
	Year          int    `json:"year" gorm:"type:int(4)"`
	CategoryId    int    `json:"categoryId"`
	Description   string `json:"description" `
}
