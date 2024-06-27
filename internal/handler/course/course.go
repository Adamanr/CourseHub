package course

import (
	courseEntity "CourseHub/internal/storage/enitity/course"
	"github.com/jackc/pgx/v5"
)

type Hub struct {
	Conn *pgx.Conn
}

type Handler interface {
	GetCourseById(id int) (*courseEntity.Course, error)
	GetCourseByTitle(title string) (*courseEntity.Course, error)
	GetAllCourses() ([]courseEntity.Course, error)
	CreateCourse(course courseEntity.Course) error
	UpdateCourse(id int) error
	DeleteCourse(id int) error
	UpdateStatusCourse(id int, status courseEntity.Status) error
	UpdateAvatarImage(id int, courseImageId string) error
}

func (c *Hub) GetCourseById(id int) (*courseEntity.Course, error) {
	course := courseEntity.Course{}

	if err := course.GetCourseById(c.Conn, id); err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *Hub) GetCourseByTitle(title string) (*courseEntity.Course, error) {
	course := courseEntity.Course{}

	if err := course.GetCourseByTitle(c.Conn, title); err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *Hub) GetAllCourses() ([]courseEntity.Course, error) {
	course := courseEntity.Course{}
	return course.GetAllCourses(c.Conn)
}

func (c *Hub) CreateCourse(course *courseEntity.Course) error {
	if err := course.CreateCourse(c.Conn); err != nil {
		return err
	}

	return nil
}

func (c *Hub) UpdateCourse(id int) error {
	course := courseEntity.Course{}

	if err := course.UpdateCourse(c.Conn, id); err != nil {
		return err
	}

	return nil
}

func (c *Hub) DeleteCourse(id int) error {
	course := courseEntity.Course{}
	return course.DeleteCourse(c.Conn, id)
}

func (c *Hub) UpdateStatusCourse(id int, status courseEntity.Status) error {
	course := courseEntity.Course{}
	if err := course.UpdateStatusCourse(c.Conn, id, status); err != nil {
		return err
	}

	return nil
}

func (c *Hub) UpdateAvatarImage(id int, courseImage string, isLink bool) error {
	course := courseEntity.Course{}

	if err := course.UpdateAvatarImage(c.Conn, id, courseImage); err != nil {
		return err
	}

	return nil
}
