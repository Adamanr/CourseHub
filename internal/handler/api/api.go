package api

import (
	userController "CourseHub/internal/handler/user"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type CourseHub struct {
	Conn *pgx.Conn
}

func (c CourseHub) Pong(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	w.Write([]byte("Ping"))
}

func NewCourseHub(db *pgx.Conn) *CourseHub {
	return &CourseHub{
		Conn: db,
	}
}

func (c CourseHub) PostCourseImage(w http.ResponseWriter, r *http.Request) {

}

func (c CourseHub) PostCourseNew(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) DeleteCourseId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetCourse(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	w.Write([]byte("Hello"))
}

func (c CourseHub) PutCourseId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetCourses(w http.ResponseWriter, r *http.Request, params GetCoursesParams) {
	//TODO implement me
	w.Write([]byte("Hello"))
}

func (c CourseHub) GetPlaylists(w http.ResponseWriter, r *http.Request, params GetPlaylistsParams) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PostPlaylistsImage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PostPlaylistsNew(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) DeletePlaylistsId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetPlaylistsId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PutPlaylistsId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) DeleteUserId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PutUserId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	users, err := userController.GetUser(c.Conn)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, users)
}

func (c CourseHub) PostUsersAvatar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PostUsersSignIn(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) PostUsersSignUp(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ ServerInterface = (*CourseHub)(nil)

// Напиши функцию которая будет красиво отдавать json вместе с данными и кодом статуса

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
