package handlers

import (
	filmsdto "backendtask/dto/films"
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

var path_file = "http://localhost:5000/uploads/"

type handlerFilm struct { //membuat struct handler dengan isi sebagai berikut
	FilmRepository repositories.FilmRepository // field UserRepository berisikan package dari repositories dan memanggil si interface UserRepository dari package repositories
}

func HandlersFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) CariFilm(c echo.Context) error {

	films, err := h.FilmRepository.CariFilm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, films)
}

func (h *handlerFilm) DapatFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.DapatFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResFilm(film)})
}

func (h *handlerFilm) MembuatFilm(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	categoryId, _ := strconv.Atoi(c.FormValue("categoryId"))
	year, _ := strconv.Atoi(c.FormValue("year"))

	meminta := filmsdto.CreatedFilmRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		Year:          year,
		CategoryId:    categoryId,
		Description:   c.FormValue("description"),
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

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "dumbflix"})

	if err != nil {
		fmt.Println(err.Error())
	}

	film := models.Film{
		Title:         meminta.Title,
		ThumbnailFilm: resp.SecureURL,
		Year:          meminta.Year,
		CategoryId:    meminta.CategoryId,
		Description:   meminta.Description,
	}

	data, err := h.FilmRepository.MembuatFilm(film)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResFilm(data)})
}

func (h *handlerFilm) UpdateFilm(c echo.Context) error {
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

	categoryId, _ := strconv.Atoi(c.FormValue("categoryId"))
	year, _ := strconv.Atoi(c.FormValue("year"))

	meminta := filmsdto.CreatedFilmRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: resp.SecureURL,
		Year:          year,
		CategoryId:    categoryId,
		Description:   c.FormValue("description"),
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.FilmRepository.DapatFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if meminta.Title != "" {
		user.Title = meminta.Title
	}
	if meminta.ThumbnailFilm != "" {
		user.ThumbnailFilm = meminta.ThumbnailFilm
	}
	if meminta.CategoryId != 0 {
		user.CategoryId = meminta.CategoryId
	}
	if meminta.Year != 0 {
		user.Year = meminta.Year
	}
	if meminta.Description != "" {
		user.Description = meminta.Description
	}
	dataCategory, _ := h.FilmRepository.DapatCatId(user.CategoryId)

	data, err := h.FilmRepository.UpdateFilm(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data.Category = dataCategory
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResFilm(data)})
}

func (h *handlerFilm) HapusFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.DapatFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FilmRepository.HapusFilm(film, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertDeleteFilm(data)})
}

func convertResFilm(u models.Film) filmsdto.FilmResponse {
	return filmsdto.FilmResponse{
		Id:            u.Id,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		Year:          u.Year,
		Category:      u.Category,
		Description:   u.Description,
	}
}
func convertDeleteFilm(u models.Film) filmsdto.FilmDeleteResponse {
	return filmsdto.FilmDeleteResponse{
		Id: u.Id,
	}
}
