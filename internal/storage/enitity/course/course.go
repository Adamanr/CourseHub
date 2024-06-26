package course

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type Status string

const (
	Closed  Status = "closed"
	Deleted Status = "deleted"
	Expired Status = "expired"
	Open    Status = "open"
)

type Course struct {
	Id            int       `json:"id,omitempty"`
	AuthorId      int       `json:"author_id"`
	CourseImageId string    `json:"course_image_id,omitempty"`
	Description   string    `json:"description"`
	EndDate       time.Time `json:"end_date,omitempty"`
	Price         float32   `json:"price"`
	StartDate     time.Time `json:"start_date,omitempty"`
	Status        Status    `json:"status,omitempty"`
	Title         string    `json:"title"`
	Category      []string  `json:"category,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type ICourse interface {
	GetCourseById(conn *pgx.Conn, id int) error
	GetCourseByTitle(conn *pgx.Conn, title string) error
	GetAllCourses(conn *pgx.Conn) ([]Course, error)
	CreateCourse(conn *pgx.Conn) error
	UpdateCourse(conn *pgx.Conn, id int) error
	DeleteCourse(conn *pgx.Conn, id int) error
	UpdateStatusCourse(conn *pgx.Conn, id int, status Status) error
	UpdateAvatarImage(conn *pgx.Conn, id int, courseImageId string) error
}

var (
	_ ICourse = (*Course)(nil)
)

func (c *Course) GetCourseById(conn *pgx.Conn, id int) error {
	rows := conn.QueryRow(context.Background(), "SELECT  * FROM courses WHERE id = $1")

	if err := rows.Scan(&c.Id, &c.AuthorId, &c.CourseImageId, &c.Title, &c.Description, &c.StartDate, &c.EndDate, &c.Category, &c.CreatedAt, &c.UpdatedAt, &c.Price, &c.Status); err != nil {
		slog.Error("Error querying users: %s", err)
		return err
	}

	return nil
}

func (c *Course) GetCourseByTitle(conn *pgx.Conn, title string) error {
	rows := conn.QueryRow(context.Background(), "SELECT  * FROM courses WHERE title = $1")

	if err := rows.Scan(&c.Id, &c.AuthorId, &c.CourseImageId, &c.Title, &c.Description, &c.StartDate, &c.EndDate, &c.Category, &c.CreatedAt, &c.UpdatedAt, &c.Price, &c.Status); err != nil {
		slog.Error("Error querying users: %s", err)
		return err
	}

	return nil
}

func (c *Course) GetAllCourses(conn *pgx.Conn) ([]Course, error) {
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

func (c *Course) CreateCourse(conn *pgx.Conn) error {
	query := `INSERT INTO courses(author_id, course_image_id, title, description, start_date, end_date, category, created_at, updated_at, price, status)
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	rows, err := conn.Exec(context.Background(), query, &c.AuthorId, &c.CourseImageId, &c.Title, &c.Description, &c.StartDate, &c.EndDate, &c.Category, time.Now(), time.Now(), &c.Price, &c.Status)
	if err != nil {
		slog.Error("Error creating course: %s", err)
		return err
	}

	slog.Info("New course created", rows.String())

	return nil
}

func (c *Course) UpdateCourse(conn *pgx.Conn, id int) error {
	query := `UPDATE courses SET title = $1, description = $2, start_date = $3, end_date = $4, category = $5, updated_at = $6, price = $7 WHERE id  =  $9`

	rows, err := conn.Exec(context.Background(), query, &c.Title, &c.Description, &c.StartDate, &c.EndDate, &c.Category, time.Now(), &c.Price, &id)
	if err != nil {
		slog.Error("Error updating course: %s", err)
		return err
	}

	slog.Info("Course updated", rows.String())

	return nil
}

func (c Course) DeleteCourse(conn *pgx.Conn, id int) error {
	query := `DELETE FROM courses WHERE id  =  $1`

	rows, err := conn.Exec(context.Background(), query, &id)
	if err != nil {
		slog.Error("Error deleting course: %s", err)
		return err
	}

	slog.Info("Course deleted", rows.String())

	return nil
}

func (c *Course) UpdateStatusCourse(conn *pgx.Conn, id int, status Status) error {
	query := `UPDATE courses SET status  =  $1 WHERE id   =   $2`

	rows, err := conn.Exec(context.Background(), query, &status, &id)
	if err != nil {
		slog.Error("Error updating course: %s", err)
		return err
	}

	slog.Info("Course status updated", rows.String())

	return nil
}

func (c Course) UpdateAvatarImage(conn *pgx.Conn, id int, courseImageId string) error {
	query := `UPDATE courses SET course_image_id  =  $1 WHERE id  =  $2`

	rows, err := conn.Exec(context.Background(), query, &courseImageId, &id)
	if err != nil {
		slog.Error("Error updating course: %s", err)
		return err
	}

	slog.Info("Course avatar updated", rows.String())

	return nil
}
