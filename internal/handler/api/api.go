package api

import (
	courseController "CourseHub/internal/handler/course"
	playlistController "CourseHub/internal/handler/playlist"
	userController "CourseHub/internal/handler/user"
	courseEntity "CourseHub/internal/storage/enitity/course"
	playlistEntity "CourseHub/internal/storage/enitity/playlist"
	userEntity "CourseHub/internal/storage/enitity/user"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"io"
	"log/slog"
	"net/http"
)

type CourseHub struct {
	Conn *pgx.Conn
}

func NewCourseHub(db *pgx.Conn) *CourseHub {
	return &CourseHub{
		Conn: db,
	}
}

var _ ServerInterface = (*CourseHub)(nil)

func (c CourseHub) DeletePlaylist(w http.ResponseWriter, r *http.Request, id int) {
	cHub := playlistController.Hub(c)

	if err := cHub.DeletePlaylist(id); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c CourseHub) GetPlaylistById(w http.ResponseWriter, r *http.Request, id int) {
	cHub := playlistController.Hub(c)

	user, err := cHub.GetPlaylistById(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, user)
}

func (c CourseHub) UpdatePlaylist(w http.ResponseWriter, r *http.Request, id int) {
	cHub := playlistController.Hub(c)

	var playlist playlistEntity.Playlist
	if err := GetObject(r, &playlist); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	user, err := cHub.UpdatePlaylist(playlist, id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, user)
}

func (c CourseHub) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	cHub := playlistController.Hub(c)

	if err := cHub.DeletePlaylist(id); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, nil)
}

func (c CourseHub) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	cHub := userController.Hub(c)

	user, err := cHub.GetUserById(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, user)
}

func (c CourseHub) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CourseHub) UploadCourseImage(w http.ResponseWriter, r *http.Request, user_id int) {
	cHub := courseController.Hub(c)

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
	cHub := courseController.Hub(c)

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
	cHub := courseController.Hub(c)

	if err := cHub.DeleteCourse(id); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (c CourseHub) GetCourseById(w http.ResponseWriter, r *http.Request, id int) {
	cHub := courseController.Hub(c)

	course, err := cHub.GetCourseById(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, course)
}

func (c CourseHub) UpdateCourse(w http.ResponseWriter, r *http.Request, id int) {
	cHub := courseController.Hub(c)

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

func (c CourseHub) UploadPlaylistImage(w http.ResponseWriter, r *http.Request, playlistId int) {
	cHub := playlistController.Hub(c)

	var image ImageStruct
	if unmErr := GetObject(r, &image); unmErr != nil {
		WriteJson(w, http.StatusBadRequest, unmErr.Error)
		return
	}

	images := image.ID
	if image.URL != "" {
		images = image.URL
	}

	playlist, createErr := cHub.UploadImage(playlistId, images)
	if createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, playlist)
}

func (c CourseHub) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	cHub := playlistController.Hub(c)

	var playlist playlistEntity.Playlist
	if err := GetObject(r, &playlist); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	newPlaylist, createErr := cHub.CreatePlaylist(playlist)
	if createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, newPlaylist)
}

func (c CourseHub) UserSignIn(w http.ResponseWriter, r *http.Request) {
	cHub := userController.Hub(c)

	var user userEntity.User
	if err := GetObject(r, &user); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	users, createErr := cHub.UserSignIn(user.Email, user.Password)
	//Tsdfsd
	if createErr != nil {

		WriteJson(w, http.StatusBadRequest, createErr.Error)

		return
	}

	WriteJson(w, http.StatusOK, users)
}

func (c CourseHub) UserSignUp(w http.ResponseWriter, r *http.Request) {
	cHub := userController.Hub(c)

	var user userEntity.User
	if err := GetObject(r, &user); err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	if createErr := cHub.UserSignUp(&user); createErr != nil {
		WriteJson(w, http.StatusBadRequest, createErr.Error)
		return
	}

	WriteJson(w, http.StatusOK, user)
}

func (c CourseHub) Pong(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, http.StatusOK, "Ping")
}

func (c CourseHub) GetCourses(w http.ResponseWriter, r *http.Request, params GetCoursesParams) {
	course := courseController.Hub(c)

	courses, err := course.GetAllCourses()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, courses)
}

func (c CourseHub) GetPlaylists(w http.ResponseWriter, r *http.Request, params GetPlaylistsParams) {
	cHub := playlistController.Hub(c)

	playlists, err := cHub.GetPlaylists()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, playlists)
}

func (c CourseHub) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	cHub := userController.Hub(c)

	users, err := cHub.GetUsers()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error)
		return
	}

	WriteJson(w, http.StatusOK, users)
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Error writing JSON to response")
		panic(err)
	}
}

func GetObject[T any](r *http.Request, obj *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body")
		return err
	}

	if unmErr := json.Unmarshal(body, &obj); unmErr != nil {
		slog.Error("Error reading JSON from request")
		return err
	}

	return nil
}
