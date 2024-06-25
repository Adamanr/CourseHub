package api

import (
	userController "CourseHub/internal/router/user"
	"encoding/json"
	"net/http"
)

type CourseHub struct{}

func NewCourseHub() *CourseHub {
	return &CourseHub{}
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
	panic("implement me")
}

func (c CourseHub) PutCourseId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) ListCourses(w http.ResponseWriter, r *http.Request, params ListCoursesParams) {
	//TODO implement me
	panic("implement me")
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
	user, err := userController.GetUsers()

	data, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
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
