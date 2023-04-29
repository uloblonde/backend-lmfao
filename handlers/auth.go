package handlers

import (
	authdto "backendtask/dto/auth"
	dto "backendtask/dto/result"
	"backendtask/models"
	"backendtask/pkg/bcrypt"
	jwtToken "backendtask/pkg/jwt"
	"backendtask/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}
func (h *handlerAuth) Login(c echo.Context) error {
	meminta := new(authdto.LoginRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    meminta.Email,
		Password: meminta.Password,
	}

	// checking user email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// checking user password
	isValid := bcrypt.CheckPassHash(meminta.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Email: user.Email,
		Token: token,
		Role:  user.Role,
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: loginResponse})
}

func (h *handlerAuth) Register(c echo.Context) error {
	meminta := new(authdto.AuthRequest)
	if err := c.Bind(meminta); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashingPassword(meminta.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    meminta.Email,
		Password: password,
		FullName: meminta.FullName,
		Gender:   meminta.Gender,
		Phone:    meminta.Phone,
		Address:  meminta.Address,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	claims := jwt.MapClaims{}
	claims["id"] = data.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token, erGenerate := jwtToken.GenerateToken(&claims)

	if erGenerate != nil {
		log.Println(erGenerate)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	registertoken := authdto.AuthResponse{
		Email: user.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Data: registertoken})

}
func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: user})
}
