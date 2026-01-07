package handler

import (
	"net/http"

	"HW_3/internal/model"
	"HW_3/internal/storage"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")

	student, err := h.storage.GetStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   student,
	})
}

func (h *Handler) GetAllSchedule(c echo.Context) error {
	results, err := h.storage.GetAllSchedule(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}

func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")

	results, err := h.storage.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}
