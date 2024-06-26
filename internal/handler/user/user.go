package user

import (
	userEntity "CourseHub/internal/storage/enitity/user"
	"github.com/jackc/pgx/v5"
)

type CourseHub struct {
	Conn *pgx.Conn
}

func GetUser(conn *pgx.Conn) ([]userEntity.User, error) {
	users, err := userEntity.GetUsers(conn)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *CourseHub) UserSignIn(email, password string) (*userEntity.User, error) {
	user := userEntity.User{}

	if err := user.UserSignIn(u.Conn, email, password); err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *CourseHub) UserSignUp(user *userEntity.User) error {
	if err := user.UserSignUp(u.Conn); err != nil {
		return err
	}

	return nil
}
