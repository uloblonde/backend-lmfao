package handlers

import (
	episdto "backendtask/dto/episode"
	dto "backendtask/dto/result"
	"backendtask/models"
	"backendtask/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerEpi struct { //membuat struct handler dengan isi sebagai berikut
	EpiRepository repositories.EpiRepository // field UserRepository berisikan package dari repositories dan memanggil si interface UserRepository dari package repositories
}

func HandlersEpi(EpiRepository repositories.EpiRepository) *handlerEpi {
	return &handlerEpi{EpiRepository}
}

func (h *handlerEpi) CariEpi(c echo.Context) error {
	films, err := h.EpiRepository.CariEpi()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, films)
}
func (h *handlerEpi) CariEpiByFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Episode, err := h.EpiRepository.CariEpiByFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: Episode})
}

func (h *handlerEpi) DapatEpiByFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ide, _ := strconv.Atoi(c.Param("ide"))

	Episode, err := h.EpiRepository.DapatEpiByFilm(id, ide)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: Episode})
}

func (h *handlerEpi) DapatEpi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.EpiRepository.DapatEpi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: film})
}

func (h *handlerEpi) MembuatEpi(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	filmId, _ := strconv.Atoi(c.FormValue("filmId"))

	meminta := episdto.CreateEpisodeRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		LinkFilm:      c.FormValue("linkFilm"),
		FilmId:        filmId,
	}

	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "dumbflix-img"})

	if err != nil {
		fmt.Println(err.Error())
	}

	epi := models.Episode{
		Title:         meminta.Title,
		ThumbnailFilm: resp.SecureURL,
		LinkFilm:      meminta.LinkFilm,
		FilmId:        meminta.FilmId,
	}
	data, err := h.EpiRepository.MembuatEpi(epi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerEpi) UpdateEpi(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "dumbflix-img"})

	if err != nil {
		fmt.Println(err.Error())
	}

	filmId, _ := strconv.Atoi(c.FormValue("filmId"))

	meminta := episdto.UpdateEpisodeRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: resp.SecureURL,
		LinkFilm:      c.FormValue("linkFilm"),
		FilmId:        filmId,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.EpiRepository.DapatEpi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if meminta.Title != "" {
		user.Title = meminta.Title
	}
	if meminta.ThumbnailFilm != "" {
		user.ThumbnailFilm = meminta.ThumbnailFilm
	}
	if meminta.LinkFilm != "" {
		user.LinkFilm = meminta.LinkFilm
	}

	data, err := h.EpiRepository.UpdateEpi(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerEpi) HapusEpi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.EpiRepository.DapatEpi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.EpiRepository.HapusEpi(film, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func convertResEpi(u models.Episode) episdto.EpisodeResponse {
	return episdto.EpisodeResponse{
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		LinkFilm:      u.LinkFilm,
	}
}
