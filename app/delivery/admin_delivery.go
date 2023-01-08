package delivery

import (
	"auth/domain"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type adminDelivery struct{}

var validate = validator.New()

func NewAdminDelivery(router fiber.Router) {
	handler := &adminDelivery{}
	router.Get("/admin", handler.Auth)
	router.Post("/admin", handler.Login)
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
	return ctx.JSON(response)
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

	secret := os.Getenv("SECRET")
	claims := jwt.MapClaims{
		"name": "jack",
		"exp":  time.Now().UTC().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return a.sendLoginResponse(ctx, http.StatusOK, err.Error(), nil)
	}

	response := a.sendLoginResponse(ctx, http.StatusOK, "Get token", &signedToken)
	return response
}

func (a *adminDelivery) Auth(ctx *fiber.Ctx) error {
	// token := ctx.GetReqHeaders()
	// fmt.Println(token)

	return ctx.Status(http.StatusUnauthorized).JSON(http.StatusText(http.StatusUnauthorized))
}
