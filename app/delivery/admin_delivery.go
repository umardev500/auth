package delivery

import (
	"auth/domain"
	"auth/helper"
	"auth/middleware"
	"auth/pb"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type adminDelivery struct {
	usecase domain.AdminUsecase
}

var validate = validator.New()

func NewAdminDelivery(router fiber.Router, storage fiber.Storage, usecase domain.AdminUsecase) {
	handler := &adminDelivery{
		usecase: usecase,
	}

	loginMaxRate, _ := strconv.Atoi(os.Getenv("LOGIN_MAX_REQ"))
	loginExpiresRate, _ := strconv.Atoi(os.Getenv("LOGIN_LIMITER_EXPIRATION_TIME"))
	limiter := middleware.RateLimiter("page-id", "192.199.9.8", storage, int64(loginMaxRate), int64(loginExpiresRate))
	router.Post("/admin", limiter, handler.Login)
	router.Get("/admin", middleware.JwtMiddleware(), handler.Auth)
}

func (a *adminDelivery) sendLoginResponse(ctx *fiber.Ctx, statusCode int, message string, token *string) error {
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
			Name:     "membership-token",
			Value:    tokenValue,
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Duration(expired) * time.Second),
		})
	}

	return ctx.JSON(response)
}

func (a *adminDelivery) createJWT(data *pb.UserLoginResponse) (token string, err error) {
	secret := os.Getenv("SECRET")
	expirationTime, _ := strconv.Atoi(os.Getenv("LOGIN_EXPIRATION_TIME"))
	claims := jwt.MapClaims{
		"user_id": data.Payload.UserId,
		"user":    data.Payload.User,
		"level":   data.Payload.Level,
		"exp":     time.Now().Add(time.Duration(expirationTime) * time.Second).Unix(),
	}

	token, err = helper.CreateJWT(secret, claims)

	return
}

func (a *adminDelivery) Login(ctx *fiber.Ctx) error {
	var req domain.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := domain.LoginResponsePayload{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		return ctx.JSON(response)
	}

	if err := validate.Struct(&req); err != nil {
		return a.sendLoginResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	// login
	reqContext := ctx.Context()
	payload := &pb.UserLoginRequest{
		User: req.Username,
		Pass: req.Password,
	}

	resp, err := a.usecase.Login(reqContext, payload)
	if err != nil {
		return a.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	if resp.IsEmpty {
		return a.sendLoginResponse(ctx, http.StatusNotFound, "Account not found", nil)
	}

	signedToken, err := a.createJWT(resp)
	if err != nil {
		return a.sendLoginResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	response := a.sendLoginResponse(ctx, http.StatusOK, "Login", &signedToken)
	return response
}

func (a *adminDelivery) Auth(ctx *fiber.Ctx) error {
	return ctx.JSON("OK")
}
