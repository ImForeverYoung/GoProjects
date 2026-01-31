package handler

import (
	"HW_5/internal/model"
	"HW_5/internal/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// Register handles user registration
func (h *AuthHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	// 1. Parse (Bind) JSON from request body to struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	// 2. Call Usecase to register user
	id, err := h.usecase.Register(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.ServerResponse{
		Status: "success",
		Data:   map[string]int{"id": id},
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	token, err := h.usecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ServerResponse{Status: "error", Message: "Invalid email or password"})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   model.AuthResponse{Token: token},
	})
}

// GetMe returns the current user's info based on the JWT token
func (h *AuthHandler) GetMe(c echo.Context) error {
	// Access the user claims set by the middleware
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// In a real app, you might want to fetch fresh details from DB using claims["user_id"]
	// For now, we return what's in the token
	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   claims,
	})
}
