package handlers

import (
	catsdto "backendtask/dto/category"
	dto "backendtask/dto/result"
	"backendtask/models"
	"backendtask/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCat struct { //membuat struct handler dengan isi sebagai berikut
	CatRepository repositories.CatRepository // field CatRepository berisikan package dari repositories dan memanggil si interface CatRepository dari package repositories
}

func HandlersCat(CatRepository repositories.CatRepository) *handlerCat {
	return &handlerCat{CatRepository}
}

func (h *handlerCat) CariCat(c echo.Context) error {
	cats, err := h.CatRepository.CariCat()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, cats)
}
func (h *handlerCat) DapatCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cat, err := h.CatRepository.DapatCat(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: cat})
}

func (h *handlerCat) MembuatCat(c echo.Context) error {
	meminta := new(catsdto.CreateCatRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	user := models.Category{
		Id:   meminta.Id,
		Name: meminta.Name,
	}
	data, err := h.CatRepository.MembuatCat(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCat) UpdateCat(c echo.Context) error {
	meminta := new(catsdto.UpdateCatRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.CatRepository.DapatCat(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if meminta.Id != 0 {
		user.Id = meminta.Id
	}
	if meminta.Name != "" {
		user.Name = meminta.Name
	}

	data, err := h.CatRepository.UpdateCat(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCat) HapusCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.CatRepository.DapatCat(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CatRepository.HapusCat(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func convertResCat(u models.Category) catsdto.CatResponse {
	return catsdto.CatResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}
