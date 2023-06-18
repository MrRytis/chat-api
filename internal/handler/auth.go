package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/model/response"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/service/authService"
	"github.com/MrRytis/chat-api/internal/service/userService"
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
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/register [post]
func Register(c *fiber.Ctx) error {
	req := new(request.Register)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	hashedPassword, err := authService.HashPassword(req.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user := userService.BuildUser(req.Email, hashedPassword, req.Name)
	repository.SaveUser(user)

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
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	req := new(request.Login)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	user := repository.FindUserByEmail(req.Email)
	if user == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email or password is incorrect")
	}

	if authService.CheckUserPassword(req.Password, user.Password) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email or password is incorrect")
	}

	return c.JSON(response.Auth{
		AccessToken:  authService.CreateAccessToken(*user),
		RefreshToken: authService.CreateRefreshToken(*user),
		ExpiresAt:    time.Now().Add(authService.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  logout user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req body request.Logout true "logout"
// @Success      204  {object}  nil
// @Failure      400  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/logout [post]
func Logout(c *fiber.Ctx) error {
	req := new(request.Logout)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	authService.BlackListToken(c.Locals("jwt").(string), c.Locals("expiresAt").(int64))
	authService.ExpireRefreshToken(c.Locals("userId").(uint), req.RefreshToken)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// Refresh godoc
// @Summary      Refresh auth token
// @Description  refreshes auth token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        req body request.Refresh true "refresh"
// @Success      200  {object}  response.Auth
// @Failure      400  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/auth/refresh [post]
func Refresh(c *fiber.Ctx) error {
	req := new(request.Refresh)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	token, err := authService.RefreshToken(req.RefreshToken, req.AccessToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	return c.JSON(response.Auth{
		AccessToken:  token,
		RefreshToken: req.RefreshToken,
		ExpiresAt:    time.Now().Add(authService.AccessTokenJwtExpDuration).Format(time.RFC3339),
	})
}
