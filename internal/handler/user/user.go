package user

import (
	userEntity "CourseHub/internal/storage/enitity/user"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type Hub struct {
	Conn *pgx.Conn
}

type Handler interface {
	GetUsers() ([]userEntity.User, error)
	GetUserById(id int) (*userEntity.User, error)
	GetUserByLogin(login string) (*userEntity.User, error)
	GetUserByEmail(email string) (*userEntity.User, error)
	UpdateUser(user *userEntity.User) error
	DeleteUser(id int) error
	ForgetPassword(email string) error
	GetAvatar(id int) (string, error)
	GetSubscribers(id int) ([]int, error)
	GetFollowers(id int) ([]int, error)
	UserSignUp(user *userEntity.User) error
	UserSignIn(email, password string) (*userEntity.User, error)
}

var _ Handler = (*Hub)(nil)

func (ch *Hub) GetUserByLogin(login string) (*userEntity.User, error) {
	user := &userEntity.User{}
	if err := user.GetUserByLogin(ch.Conn, login); err != nil {
		slog.Error("User error GetUserByLogin:   ", err.Error)
		return nil, err
	}

	return user, nil
}

func (ch *Hub) GetUserByEmail(email string) (*userEntity.User, error) {
	user := &userEntity.User{}
	if err := user.GetUserByEmail(ch.Conn, email); err != nil {
		slog.Error("User error GetUserByEmail:   ", err.Error)
		return nil, err
	}

	return user, nil
}

func (ch *Hub) UpdateUser(user *userEntity.User) error {
	//TODO implement me
	panic("implement me")
}

func (ch *Hub) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (ch *Hub) GetAvatar(id int) (string, error) {
	user := &userEntity.User{}
	if err := user.GetAvatar(ch.Conn, id); err != nil {
		slog.Error("User error GetAvatar:  ", err.Error)
		return "", err
	}

	return *user.AvatarId, nil
}

func (ch *Hub) GetSubscribers(id int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (ch *Hub) GetFollowers(id int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (ch *Hub) GetUsers() ([]userEntity.User, error) {
	user := &userEntity.User{}
	return user.GetUsers(ch.Conn)
}

func (ch *Hub) ForgetPassword(email string) error {
	user := &userEntity.User{}

	if err := user.ForgetPassword(ch.Conn, email); err != nil {
		slog.Error("User error ForgetPassword: ", err.Error)
		return err
	}

	return nil
}

func (ch *Hub) GetUserById(id int) (*userEntity.User, error) {
	user := &userEntity.User{}

	if err := user.GetUserById(ch.Conn, id); err != nil {
		slog.Error("User error GetUserById: ", err.Error)
		return nil, err
	}

	return user, nil
}

func (ch *Hub) UserSignIn(email, password string) (*userEntity.User, error) {
	user := userEntity.User{}

	if err := user.UserSignIn(ch.Conn, email, password); err != nil {
		slog.Error("User error SignIn: ", err.Error)
		return nil, err
	}

	return &user, nil
}

func (ch *Hub) UserSignUp(user *userEntity.User) error {
	if err := user.UserSignUp(ch.Conn); err != nil {
		slog.Error("User error SignUp: ", err.Error)
		return err
	}

	return nil
}
