package delivery

import (
	"auth/domain"
	"auth/helper"
	"auth/middleware"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type adminDelivery struct{}

var validate = validator.New()

func NewAdminDelivery(router fiber.Router, storage fiber.Storage) {
	handler := &adminDelivery{}

	router.Post("/admin", middleware.RateLimiter(storage), handler.Login)
	router.Get("/admin", middleware.JwtMiddleware(), handler.Auth)
}

func (a *adminDelivery) sendLoginResponse(ctx *fiber.Ctx, statusCode int, message string, token *string) error {
	var tokenValue string
	if token != nil {
		tokenValue = *token
	}

	response := domain.LoginResponse{
		StatusCode: statusCode,
		Message:    message,
		Token:      tokenValue,
	}

	expired, _ := strconv.Atoi(os.Getenv("COOKIE_EXPIRATION_TIME"))

	ctx.Cookie(&fiber.Cookie{
		Path:     "/",
		Name:     "token",
		Value:    *token,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Duration(expired) * time.Second),
	})

	return ctx.JSON(response)
}

func (a *adminDelivery) createJWT() (token string, err error) {
	secret := os.Getenv("SECRET")
	expirationTime, _ := strconv.Atoi(os.Getenv("LOGIN_EXPIRATION_TIME"))
	claims := jwt.MapClaims{
		"name": "jack",
		"exp":  time.Now().Add(time.Duration(expirationTime) * time.Second).Unix(),
	}

	token, err = helper.CreateJWT(secret, claims)

	return
}

func (a *adminDelivery) Login(ctx *fiber.Ctx) error {
	var req domain.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := domain.LoginResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		return ctx.JSON(response)
	}

	if err := validate.Struct(&req); err != nil {
		return a.sendLoginResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	signedToken, err := a.createJWT()
	if err != nil {
		return a.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	response := a.sendLoginResponse(ctx, http.StatusOK, "Get token", &signedToken)
	return response
}

func (a *adminDelivery) Auth(ctx *fiber.Ctx) error {
	return ctx.JSON("OK")
}
