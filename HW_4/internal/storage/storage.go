package storage

import (
	"HW_4/internal/model"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

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

	rows, err := s.conn.Query(ctx, `select group_name, subject_name, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id 
			join subjects on schedules.subject_id = subjects.subject_id`)
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

	rows, err := s.conn.Query(ctx, `select group_name, subject_name, time_slot 
			from schedules join groups on schedules.group_id = groups.group_id
			join subjects on schedules.subject_id = subjects.subject_id
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

func (s *Storage) MarkAttendance(ctx context.Context, request model.Attendance) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	queryI := `INSERT INTO attendance (student_id, subject_id, visit_day, visited)
VALUES ($1, $2, $3, $4) RETURNING id;`

	row := s.conn.QueryRow(ctx, queryI,
		request.StudentID,
		request.SubjectID,
		request.VisitDay,
		request.Visited,
	)
	err = row.Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return 0, err
	}

	return id, nil

}

func (s *Storage) GetAttendanceBySubjectId(ctx context.Context, id string) ([]model.Attendance, error) {
	results := []model.Attendance{}

	rows, err := s.conn.Query(ctx, `select student_id, visit_day::text, visited
			 from attendance
			 where subject_id=$1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var studentId int
		var visitedDay string
		var visited bool

		rows.Scan(&studentId, &visitedDay, &visited)
		idInt, _ := strconv.Atoi(id)
		attendance := model.Attendance{
			StudentID: studentId,
			SubjectID: idInt,
			VisitDay:  visitedDay,
			Visited:   visited,
		}
		results = append(results, attendance)
	}
	return results, nil
}

func (s *Storage) GetAttendanceByStudentId(ctx context.Context, id string) ([]model.Attendance, error) {
	results := []model.Attendance{}

	rows, err := s.conn.Query(ctx, `select subject_id, visit_day::text, visited
			 from attendance
			 where student_id=$1`, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subjectId int
		var visitedDay string
		var visited bool

		rows.Scan(&subjectId, &visitedDay, &visited)
		idInt, _ := strconv.Atoi(id)
		attendance := model.Attendance{
			StudentID: idInt,
			SubjectID: subjectId,
			VisitDay:  visitedDay,
			Visited:   visited,
		}
		results = append(results, attendance)
	}
	return results, nil
}
