package handlers

import (
	dto "backendtask/dto/result"
	usersdto "backendtask/dto/users"
	"backendtask/models"
	"backendtask/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handler struct { //membuat struct handler dengan isi sebagai berikut
	UserRepository repositories.UserRepository // field UserRepository berisikan package dari repositories dan memanggil si interface UserRepository dari package repositories
}

func HandlersUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) CariUser(c echo.Context) error {
	users, err := h.UserRepository.CariUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
func (h *handler) CariProfile(c echo.Context) error {
	userId := c.Get("userLogin").(jwt.MapClaims)["id"].(float64)

	var profile models.User
	profile, err := h.UserRepository.DapatUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: convertResponse(profile)})
}

func (h *handler) DapatUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.DapatUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: user})
}

func (h *handler) MembuatUser(c echo.Context) error {
	meminta := new(usersdto.CreatedUserRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	user := models.User{
		Email:    meminta.Email,
		Password: meminta.Password,
		FullName: meminta.FullName,
		Gender:   meminta.Gender,
		Phone:    meminta.Phone,
		Address:  meminta.Address,
	}
	data, err := h.UserRepository.MembuatUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handler) UpdateUser(c echo.Context) error {
	meminta := new(usersdto.UpdateUserRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.DapatUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if meminta.Email != "" {
		user.Email = meminta.Email
	}
	if meminta.Password != "" {
		user.Password = meminta.Password
	}
	if meminta.FullName != "" {
		user.FullName = meminta.FullName
	}
	if meminta.Gender != "" {
		user.Gender = meminta.Gender
	}
	if meminta.Phone != "" {
		user.Phone = meminta.Phone
	}
	if meminta.Address != "" {
		user.Address = meminta.Address
	}

	data, err := h.UserRepository.UpdateUser(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handler) HapusUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.DapatUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.HapusUser(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		Id:       u.Id,
		FullName: u.FullName,
		Email:    u.Email,
		Password: u.Password,
		Gender:   u.Gender,
		Phone:    u.Phone,
		Address:  u.Address,
	}
}

// type Login struct {
// 	Id       string `json:"id"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	FullName string `json:"fullName"`
// 	Gender   string `json:"gender"`
// 	Phone    string `json:"phone"`
// 	Address  string `json:"address"`
// }

// var iniLogin = []Login{
// 	{
// 		Id:       "1",
// 		FullName: "Iis Is",
// 		Email:    "iis@gmail.com",
// 		Password: "lovespiderman",
// 		Gender:   "male",
// 		Phone:    "083896831233",
// 		Address:  "Jln. Marvel Universe, RT.21 RW.69",
// 	},
// }

// func MencariLogin(c echo.Context) error {
// 	c.Response().Header().Set("Content-Type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(iniLogin)
// }
// func MencariLgnId(c echo.Context) error {

// 	c.Response().Header().Set("content-type", "application/json")
// 	id := c.Param("id")    //inisialisasi param id
// 	var keData Login       //inisialisasi dari struct (Login)
// 	var IniKondisi = false //inisialisasi boolean

// 	for _, data := range iniLogin { //perulangan dari variable data di dalam range array of Object iniLogin
// 		if id == data.Id { //pengkondisian  "jika id(parameter) == variabel data.Id(array of Object) bernilai true"
// 			IniKondisi = true //kondisi variabel Ini Kondisi true
// 			keData = data     //struct keData adalah data dari iniLogin (array of Object)
// 		}
// 	}

// 	if !IniKondisi { // Jika tidak samadengan IniKodisi
// 		c.Response().WriteHeader(http.StatusNotFound)                           //response akan mengeluarkan status not found (404)
// 		return json.NewEncoder(c.Response()).Encode("ID : " + id + "Not Found") //Message dari Encoder
// 	}
// 	c.Response().WriteHeader(http.StatusOK)             //bila hasilnya ada true dia akan mengeluarkan status OK (200)
// 	return json.NewEncoder(c.Response()).Encode(keData) //response json Encode keData(struct)
// }

// func CreateLogin(c echo.Context) error {
// 	var data Login //inisialisasi struct Login

// 	json.NewDecoder(c.Request().Body).Decode(&data) //Decode dari jsonnya

// 	iniLogin = append(iniLogin, data) //mengappend data ke objek iniLogin

// 	c.Response().Header().Set("Content-Type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(data)
// }

// func UpdateLogin(c echo.Context) error {
// 	id := c.Param("id")
// 	var data Login //inisialisasi struct Login
// 	var iniKondisi = false

// 	json.NewDecoder(c.Request().Body).Decode(&data)

// 	for idx, data := range iniLogin {
// 		if id == data.Id {
// 			iniKondisi = true
// 			iniLogin[idx] = data
// 		}
// 	}
// 	if !iniKondisi {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID : " + id + "Not Found")
// 	}
// 	c.Response().Header().Set("content-type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(data)
// }

// func DeleteLogin(c echo.Context) error {
// 	id := c.Param("id")
// 	var iniKondisi = false
// 	var index = 0

// 	for idx, data := range iniLogin {
// 		if id == data.Id {
// 			iniKondisi = true
// 			index = idx
// 		}
// 	}

// 	if !iniKondisi {
// 		c.Response().WriteHeader(http.StatusNotFound)
// 		return json.NewEncoder(c.Response()).Encode("ID : " + id + "Not Found")
// 	}
// 	iniLogin = append(iniLogin[:index], iniLogin[index+1:]...)
// 	c.Response().Header().Set("content-type", "application/json")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return json.NewEncoder(c.Response()).Encode(iniLogin)
// }
