package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"HW_4/internal/handler"
	"HW_4/internal/storage"
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

	// Init Handler
	h := handler.NewHandler(store)

	// Init Echo
	e := echo.New()

	// Routes
	e.GET("/student/:id", h.GetStudent)
	e.GET("/all_class_schedule", h.GetAllSchedule)
	e.GET("/schedule/group/:id", h.GetGroupSchedule)
	e.POST("/attendance/subject", h.MarkAttendance)
	e.GET("/attendanceBySubjectId/:id", h.GetAttendanceBySubjectId)
	e.GET("/attendanceByStudentId/:id", h.GetAttendanceByStudentId)
	
	e.Logger.Fatal(e.Start(":1323"))
}
