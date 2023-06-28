package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/model/response"
	"github.com/MrRytis/chat-api/internal/service/authService"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Register godoc
// @Summary      Register new user
// @Description  register new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req body request.Register true "register"
// @Success      201  {object}  response.Register
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	req := new(request.Register)
	utils.ParseBodyAndValidate(c, req)

	user := authService.RegisterUser(req.Email, req.Password, req.Name)

	return c.Status(fiber.StatusCreated).JSON(response.Register{
		UserId:  user.UUID,
		Message: "User created",
	})
}

// Login godoc
// @Summary      Login user
// @Description  login user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req body request.Login true "login"
// @Success      200  {object}  response.Auth
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	req := new(request.Login)
	utils.ParseBodyAndValidate(c, req)

	user := authService.FindUserByEmailAndPassword(req.Email, req.Password)
	accessToken := authService.CreateAccessToken(user)
	refreshToken := authService.CreateRefreshToken(user)

	return c.JSON(response.Auth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(authService.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  logout user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        req body request.Logout true "logout"
// @Success      204  {object}  nil
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/logout [post]
func Logout(c *fiber.Ctx) error {
	req := new(request.Logout)
	utils.ParseBodyAndValidate(c, req)

	authService.LogoutUser(
		c.Locals("jwt").(string),
		req.RefreshToken,
		c.Locals("userId").(uint),
		c.Locals("expiresAt").(int64),
	)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// Refresh godoc
// @Summary      Refresh auth token
// @Description  refreshes auth token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        req body request.Refresh true "refresh"
// @Success      200  {object}  response.Auth
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/refresh [post]
func Refresh(c *fiber.Ctx) error {
	req := new(request.Refresh)
	utils.ParseBodyAndValidate(c, req)

	accessToken := authService.RefreshAccessToken(req.RefreshToken, req.AccessToken)

	return c.JSON(response.Auth{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
		ExpiresAt:    time.Now().Add(authService.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}
