package handler

import (
	"net/http"
	"fmt"
	"HW_5/internal/model"
	"HW_5/internal/storage"

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


func (h *Handler) MarkAttendance(ctx echo.Context) error {
	var request model.Attendance

	if err := ctx.Bind(&request); err != nil { return ctx.JSON(http.StatusBadRequest, err) }


	id, err := h.storage.MarkAttendance(ctx.Request().Context(), request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}
	message := fmt.Sprintf("google sheet with id=%d succefully wroten in database", id)
	return ctx.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   message,
	})
}


func (h *Handler) GetAttendanceBySubjectId(c echo.Context) error {
	id := c.Param("id")

	results, err := h.storage.GetAttendanceBySubjectId(c.Request().Context(), id)
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

func (h *Handler) GetAttendanceByStudentId(c echo.Context) error {
	id := c.Param("id")
	
	results, err := h.storage.GetAttendanceByStudentId(c.Request().Context(), id)
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
