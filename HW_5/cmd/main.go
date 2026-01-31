package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"HW_5/internal/handler"
	"HW_5/internal/middleware"
	"HW_5/internal/storage"
	"HW_5/internal/usecase"
)

func main() {
	godotenv.Load()

	// Init DB
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Init Storage
	store := storage.NewStorage(conn)

	// Init Auth Components
	authUsecase := usecase.NewAuthUsecase(store)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Init Handler
	h := handler.NewHandler(store)

	// Init Echo
	e := echo.New()

	// Auth Routes
	e.POST("/api/auth/register", authHandler.Register)
	e.POST("/api/auth/login", authHandler.Login)

	// Protected Routes
	// Group "api/users" and apply custom JWT Middleware
	userGroup := e.Group("/api/users")
	userGroup.Use(middleware.JWTMiddleware)
	userGroup.GET("/me", authHandler.GetMe)

	
	e.GET("/student/:id", h.GetStudent)
	e.GET("/all_class_schedule", h.GetAllSchedule)
	e.GET("/schedule/group/:id", h.GetGroupSchedule)
	e.POST("/attendance/subject", h.MarkAttendance)
	e.GET("/attendanceBySubjectId/:id", h.GetAttendanceBySubjectId)
	e.GET("/attendanceByStudentId/:id", h.GetAttendanceByStudentId)

	e.Logger.Fatal(e.Start(":1323"))
}
