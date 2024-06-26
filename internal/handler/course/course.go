package course

import (
	courseEntity "CourseHub/internal/storage/enitity/course"
	"github.com/jackc/pgx/v5"
)

type CourseHub struct {
	Conn *pgx.Conn
}

type CourseHandler interface {
	GetCourseById(id int) (*courseEntity.Course, error)
	GetCourseByTitle(title string) (*courseEntity.Course, error)
	GetAllCourses() ([]courseEntity.Course, error)
	CreateCourse(course courseEntity.Course) error
	UpdateCourse(id int) error
	DeleteCourse(id int) error
	UpdateStatusCourse(id int, status courseEntity.Status) error
	UpdateAvatarImage(id int, courseImageId string) error
}

func (c *CourseHub) GetCourseById(id int) (*courseEntity.Course, error) {
	course := courseEntity.Course{}

	if err := course.GetCourseById(c.Conn, id); err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *CourseHub) GetCourseByTitle(title string) (*courseEntity.Course, error) {
	course := courseEntity.Course{}

	if err := course.GetCourseByTitle(c.Conn, title); err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *CourseHub) GetAllCourses() ([]courseEntity.Course, error) {
	course := courseEntity.Course{}
	return course.GetAllCourses(c.Conn)
}

func (c *CourseHub) CreateCourse(course *courseEntity.Course) error {
	if err := course.CreateCourse(c.Conn); err != nil {
		return err
	}

	return nil
}

func (c *CourseHub) UpdateCourse(id int) error {
	course := courseEntity.Course{}

	if err := course.UpdateCourse(c.Conn, id); err != nil {
		return err
	}

	return nil
}

func (c *CourseHub) DeleteCourse(id int) error {
	course := courseEntity.Course{}
	return course.DeleteCourse(c.Conn, id)
}

func (c *CourseHub) UpdateStatusCourse(id int, status courseEntity.Status) error {
	course := courseEntity.Course{}
	if err := course.UpdateStatusCourse(c.Conn, id, status); err != nil {
		return err
	}

	return nil
}

func (c *CourseHub) UpdateAvatarImage(id int, courseImage string, isLink bool) error {
	course := courseEntity.Course{}

	if err := course.UpdateAvatarImage(c.Conn, id, courseImage); err != nil {
		return err
	}

	return nil
}
