package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"HW_3/internal/model"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) GetStudent(ctx context.Context, id string) (model.StudentDetailResponse, error) {
	var fullName string
	var gender string
	var birthDate string
	var groupName string

	err := s.conn.QueryRow(ctx, `select student_name, gender, birth_date::text, group_name 
		from students join groups on students.group_id = groups.group_id
		 where students.student_id=$1`, id).Scan(&fullName, &gender, &birthDate, &groupName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return model.StudentDetailResponse{}, err
	}

	idInt, _ := strconv.Atoi(id)

	return model.StudentDetailResponse{
		ID:        idInt,
		FullName:  fullName,
		Gender:    gender,
		BirthDate: birthDate,
		GroupName: groupName,
	}, nil
}

func (s *Storage) GetAllSchedule(ctx context.Context) ([]model.ScheduleResponse, error) {
	results := []model.ScheduleResponse{}

	rows, err := s.conn.Query(ctx, `select group_name, subject, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		var subject string
		var timeSlot string

		rows.Scan(&groupName, &subject, &timeSlot)

		schedule := model.ScheduleResponse{
			GroupName: groupName,
			Subject:   subject,
			TimeSlot:  timeSlot,
		}
		results = append(results, schedule)
	}
	return results, nil
}

func (s *Storage) GetGroupSchedule(ctx context.Context, id string) ([]model.ScheduleResponse, error) {
	results := []model.ScheduleResponse{}

	rows, err := s.conn.Query(ctx, `select group_name, subject, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id
			 where groups.group_id=$1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		var subject string
		var timeSlot string

		rows.Scan(&groupName, &subject, &timeSlot)

		schedule := model.ScheduleResponse{
			GroupName: groupName,
			Subject:   subject,
			TimeSlot:  timeSlot,
		}
		results = append(results, schedule)
	}
	return results, nil
}
