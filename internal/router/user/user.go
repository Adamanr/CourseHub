package user

import (
	"CourseHub/internal/storage/enitity/user"
)

func GetUsers() (*user.User, error) {

	return &user.User{
		Id:    1,
		Login: "admin",
	}, nil
}
