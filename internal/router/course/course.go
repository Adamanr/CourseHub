package course

import (
	"CourseHub/internal/handler/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ICourseRouter interface {
	GetCourse(ch *config.CourseHub) func(w http.ResponseWriter, r *http.Request)
	CreateCourse(ch *config.CourseHub) func(w http.ResponseWriter, r *http.Request)
	UpdateCourse(ch *config.CourseHub) func(w http.ResponseWriter, r *http.Request)
	DeleteCourse(ch *config.CourseHub) func(w http.ResponseWriter, r *http.Request)
	GetCourses(ch *config.CourseHub) func(w http.ResponseWriter, r *http.Request)
}

func SetCourseRouter(r *chi.Mux) {
	r.Get("/", GetCourse)
	r.Post("/", CreateCourse)
	r.Put("/", UpdateCourse)
	r.Delete("/", DeleteCourse)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCourse"))
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateCourse"))
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateCourse"))
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteCourse"))
}
