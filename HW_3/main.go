package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type ServerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


type StudentDetailResponse struct {
	ID        int    `json:"id"`          
	FullName  string `json:"full_name"`   
	Gender    string `json:"gender"`      
	BirthDate string `json:"birth_date"`  
	GroupName string `json:"group_name"`  
}


type ScheduleResponse struct {
	GroupName  string `json:"group_name"`
	Subject string `json:"subject"`
	TimeSlot string `json:"time_slot"`
}








func main() {

	godotenv.Load()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	e := echo.New()
	e.GET("/student/:id", func(c echo.Context) error {
		
		id := c.Param("id")
		var fullName string
		var gender string
		var birthDate string

		var groupName string

		err = conn.QueryRow(context.Background(), `select student_name, gender, birth_date::text, group_name 
		from students join groups on students.group_id = groups.group_id
		 where students.student_id=$1`, id).Scan(&fullName, &gender, &birthDate, &groupName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return c.JSON(http.StatusInternalServerError, ServerResponse{
				Status:  "error",
				Message: "Internal Server Error",
			})
		}

		

		idInt, _ := strconv.Atoi(id)

		student := StudentDetailResponse{
			ID:        idInt,
			FullName:  fullName,
			Gender:    gender,
			BirthDate: birthDate,
			GroupName: groupName,
		}

		response := ServerResponse{
			Status: "success",
			Data:   student, 
		}

    	return c.JSON(http.StatusOK, response)
	})
	e.GET("/all_class_schedule", func(c echo.Context) error {
		
		results := []ScheduleResponse{}

		
		rows, err := conn.Query(context.Background(), `select group_name, subject, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
			return c.JSON(http.StatusInternalServerError, ServerResponse{
				Status:  "error",
				Message: "Internal Server Error",
			})
		}
		defer rows.Close()
		for rows.Next() { 
			var groupName string
			var subject string
			var timeSlot string

			rows.Scan(&groupName, &subject, &timeSlot)

			schedule := ScheduleResponse{
				GroupName:  groupName,
				Subject: subject,
				TimeSlot: timeSlot,
			}
			results = append(results, schedule)
		}

		

		response := ServerResponse{
			Status: "success",
			Data:   results, 
		}

    	return c.JSON(http.StatusOK, response)
	})
	e.GET("/schedule/group/:id", func(c echo.Context) error {
		//ss
		id := c.Param("id")

		results := []ScheduleResponse{}

	

		rows, err := conn.Query(context.Background(), `select group_name, subject, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id
			 where groups.group_id=$1`, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return c.JSON(http.StatusInternalServerError, ServerResponse{
				Status:  "error",
				Message: "Internal Server Error",
			})
		}
		defer rows.Close()
		for rows.Next() { 
			var groupName string
			var subject string
			var timeSlot string

			rows.Scan(&groupName, &subject, &timeSlot)

			schedule := ScheduleResponse{
				GroupName:  groupName,
				Subject: subject,
				TimeSlot: timeSlot,
			}
			results = append(results, schedule)
		}
		

		

		response := ServerResponse{
			Status: "success",
			Data:   results, 
		}

    	return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
