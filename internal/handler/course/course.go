package course

import (
	"CourseHub/internal/storage/enitity/course"
	"github.com/jackc/pgx/v5"
)

func GetCourses(conn *pgx.Conn) ([]course.Course, error) {
	courses, err := course.GetCourses(conn)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
