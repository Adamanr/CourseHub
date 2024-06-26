package user

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Login       string    `yaml:"login, omitempty"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AvatarId    *string   `json:"avatar_id,omitempty"`
	Description *string   `json:"description,omitempty" `
	CreatedAt   time.Time `json:"created_at,omitempty" `
	UpdatedAt   time.Time `json:"updated_at,omitempty" `
}

type IUser interface {
	GetUsers() ([]*User, error)
	GetUserById(id int) (*User, error)
	GetUserByLogin(login string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
	ForgetPassword(email string) error
	GetAvatar(id int) (string, error)
	GetSubscribers(id int) ([]int, error)
	GetFollowers(id int) ([]int, error)
	UserSignUp(id int) ([]*User, error)
	UserSignIn(id int) ([]*User, error)
}

func GetUsers(conn *pgx.Conn) ([]User, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		slog.Error("Error querying users: %s", err)
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if scanErr := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.AvatarId, &user.Description, &user.CreatedAt, &user.UpdatedAt); scanErr != nil {
			slog.Error("Error scanning row: %s", scanErr)
			return nil, scanErr
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetUserById(id int) (*User, error) {
	return nil, nil
}

func (u *User) GetUserByLogin(login string) (*User, error) {
	return nil, nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

func (u *User) CreateUser(conn *pgx.Conn) (*User, error) {
	return nil, nil
}

func (u *User) UpdateUser(user *User) error {
	return nil
}

func (u *User) DeleteUser(id int) error {
	return nil
}

func (u *User) ForgetPassword(email string) error {
	return nil
}

func (u *User) GetAvatar(id int) (string, error) {
	return "", nil
}

func (u *User) GetSubscribers(id int) ([]int, error) {
	return nil, nil
}

func (u *User) GetFollowers(id int) ([]int, error) {
	return nil, nil
}

func (u *User) UserSignIn(conn *pgx.Conn, email string, password string) error {
	rows := conn.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1 and password = $2", email, password)

	var user User
	if scanErr := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.AvatarId, &user.Description, &user.CreatedAt, &user.UpdatedAt); scanErr != nil {
		slog.Error("Error scanning row: %s", scanErr)
		return scanErr
	}
	return nil
}

func (u *User) UserSignUp(conn *pgx.Conn) error {

	query := `INSERT INTO users(login, email, password, created_at)
			  VALUES($1, $2, $3)`

	rows, err := conn.Exec(context.Background(), query, &u.Login, &u.Email, &u.Password)

	if err != nil {
		slog.Error("Error creating user: %s", err)
		return err
	}

	slog.Info("New user created", rows.String())

	return nil
}
