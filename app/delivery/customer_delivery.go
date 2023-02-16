package delivery

import (
	"auth/domain"
	"auth/helper"
	"auth/middleware"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type customerDelivery struct {
	usecase domain.CustomerUsecase
}

func NewCustomerDeliery(router fiber.Router, storage fiber.Storage, usecase domain.CustomerUsecase) {
	handler := &customerDelivery{
		usecase: usecase,
	}
	loginMaxRate, _ := strconv.Atoi(os.Getenv("LOGIN_MAX_REQ"))
	loginExpiresRate, _ := strconv.Atoi(os.Getenv("LOGIN_LIMITER_EXPIRATION_TIME"))
	limiter := middleware.RateLimiter("page-id", "192.199.9.8", storage, int64(loginMaxRate), int64(loginExpiresRate))
	router.Post("/login", limiter, handler.Login)
	router.Get("/auth", middleware.JwtMiddleware(), handler.Auth)
}

func (c *customerDelivery) sendLoginResponse(ctx *fiber.Ctx, statusCode int, message string, token *string) error {
	var tokenValue string
	if token != nil {
		tokenValue = *token
	}

	response := domain.LoginResponsePayload{
		StatusCode: statusCode,
		Message:    message,
		Token:      tokenValue,
	}

	expired, _ := strconv.Atoi(os.Getenv("COOKIE_EXPIRATION_TIME"))

	if token != nil {
		ctx.Cookie(&fiber.Cookie{
			Path:     "/",
			Name:     "token",
			Value:    tokenValue,
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Duration(expired) * time.Second),
		})
	}

	return ctx.JSON(response)
}

func (c *customerDelivery) createJWT(data domain.LoginResponseData) (token string, err error) {
	secret := os.Getenv("SECRET")
	expirationTime, _ := strconv.Atoi(os.Getenv("LOGIN_EXPIRATION_TIME"))
	claims := jwt.MapClaims{
		"user_id": data.UserId,
		"user":    data.User,
		"exp":     time.Now().Add(time.Duration(expirationTime) * time.Second).Unix(),
	}

	token, err = helper.CreateJWT(secret, claims)
	return
}

func (c *customerDelivery) Login(ctx *fiber.Ctx) error {

	var payload domain.LoginRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return c.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	reqCtx := ctx.Context()
	resp, err := c.usecase.Login(reqCtx, domain.LoginRequest{Username: payload.Username, Password: payload.Password})
	if err != nil {
		return c.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	if resp.IsEmpty {
		return c.sendLoginResponse(ctx, http.StatusNotFound, "Account not found", nil)
	}

	token, err := c.createJWT(resp.Payload)
	if err != nil {
		return c.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	return c.sendLoginResponse(ctx, http.StatusOK, "Login customer", &token)
}

func (c *customerDelivery) Auth(ctx *fiber.Ctx) error {
	return ctx.JSON("OK")
}
