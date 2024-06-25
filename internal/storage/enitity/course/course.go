package course

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

const (
	Closed  CourseStatus = "closed"
	Deleted CourseStatus = "deleted"
	Expired CourseStatus = "expired"
	Open    CourseStatus = "open"
)

type Course struct {
	Id            int
	AuthorId      int
	CourseImageId string
	Description   string
	EndDate       time.Time
	Price         float32
	StartDate     time.Time
	Status        CourseStatus
	Title         string
	Category      []string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CourseStatus string

func GetCourses(conn *pgx.Conn) ([]Course, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM courses")
	if err != nil {
		slog.Error("Error querying users: %s", err)
		return nil, err
	}

	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var c Course

		if scanErr := rows.Scan(&c.Id, &c.AuthorId, &c.CourseImageId, &c.Title, &c.Description, &c.StartDate, &c.EndDate, &c.Category, &c.CreatedAt, &c.UpdatedAt, &c.Price, &c.Status); scanErr != nil {
			slog.Error("Error scanning row: %s", scanErr)
			return nil, scanErr
		}

		courses = append(courses, c)
	}

	return courses, nil
}
