package api

import (
	courseController "CourseHub/internal/handler/course"
	playlistController "CourseHub/internal/handler/playlsit"
	userController "CourseHub/internal/handler/user"
	courseEntity "CourseHub/internal/storage/enitity/course"
	"CourseHub/internal/storage/enitity/course"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"io"
	"net/http"
)

type CourseHub struct {
	Conn *pgx.Conn
}

var _ ServerInterface = (*CourseHub)(nil)

func (c CourseHub) UploadCourseImage(w http.ResponseWriter, r *http.Request, user_id int) {
	cHub := courseController.CourseHub(c)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	var is ImageStruct
	if unmErr := json.Unmarshal(body, &is); unmErr != nil {
		WriteJson(w, http.StatusBadRequest, unmErr.Error)
		return
	}

	image := is.ID
	isLink := false
	if is.URL != "" {
		isLink = true
		is.ID = is.URL
	}

	if createErr := cHub.UpdateAvatarImage(user_id, image, isLink); createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, map[string]string{
		"message": "Аватарка была успешно обновлена",
	})
}

func (c CourseHub) CreateCourse(w http.ResponseWriter, r *http.Request) {
	cHub := courseController.CourseHub(c)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	var course courseEntity.Course
	if unmErr := json.Unmarshal(body, &course); unmErr != nil {
		WriteJson(w, http.StatusBadRequest, unmErr.Error)
		return
	}

	if createErr := cHub.CreateCourse(&course); createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, course)
}

func (c CourseHub) DeleteCourse(w http.ResponseWriter, r *http.Request, id int) {
	cHub := courseController.CourseHub(c)

	if err := cHub.DeleteCourse(id); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (c CourseHub) GetCourseById(w http.ResponseWriter, r *http.Request, id int) {
	cHub := courseController.CourseHub(c)

	course, err := cHub.GetCourseById(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, course)
}

func (c CourseHub) UpdateCourse(w http.ResponseWriter, r *http.Request, id int) {
	cHub := courseController.CourseHub(c)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	var course courseEntity.Course
	if unmErr := json.Unmarshal(body, &course); unmErr != nil {
		WriteJson(w, http.StatusBadRequest, unmErr.Error)
		return
	}

	if createErr := cHub.UpdateCourse(id); createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, course)
}

type ImageStruct struct {
	ID  string `json:"external_id,omitempty"`
	URL string `json:"external_url,omitempty"`
}

func (c CourseHub) UploadImage(w http.ResponseWriter, r *http.Request) {

}

func (c CourseHub) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
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
	courses, err := course.GetCourses(c.Conn)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, courses)

}

func (c CourseHub) GetPlaylists(w http.ResponseWriter, r *http.Request, params GetPlaylistsParams) {
	playlists, err := playlist.GetPlaylists(c.Conn)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, playlists)
}

func (c CourseHub) UpdatePlaylist(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UserSignIn(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UserSignUp(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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

func (c CourseHub) GetCourses(w http.ResponseWriter, r *http.Request, params GetCoursesParams) {
	course := courseController.CourseHub(c)

	courses, err := course.GetAllCourses()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, courses)
}

func (c CourseHub) GetPlaylists(w http.ResponseWriter, r *http.Request, params GetPlaylistsParams) {
	playlists, err := playlist.GetPlaylists(c.Conn)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, playlists)
}

func (c CourseHub) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	users, err := userController.GetUser(c.Conn)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, users)
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
