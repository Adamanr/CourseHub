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
	GetUsers(conn *pgx.Conn) ([]User, error)
	GetUserById(conn *pgx.Conn, id int) error
	GetUserByLogin(conn *pgx.Conn, login string) error
	GetUserByEmail(conn *pgx.Conn, email string) error
	//Оставить
	UpdateUser(conn *pgx.Conn, user *User) error
	DeleteUser(conn *pgx.Conn, id int) error
	//Оставить
	ForgetPassword(conn *pgx.Conn, email string) error
	GetAvatar(conn *pgx.Conn, id int) error
	GetSubscribers(conn *pgx.Conn, id int) ([]int, error)
	GetFollowers(conn *pgx.Conn, id int) ([]int, error)
	UserSignUp(conn *pgx.Conn) error
	UserSignIn(conn *pgx.Conn, email string, password string) error
}

var _ IUser = (*User)(nil)

func (u *User) GetUserById(conn *pgx.Conn, id int) error {
	query := `SELECT * FROM users WHERE id = $1`

	row := conn.QueryRow(context.Background(), query, id)
	if err := row.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.AvatarId, &u.Description, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByLogin(conn *pgx.Conn, login string) error {
	query := `SELECT * FROM users WHERE login = $1`

	row := conn.QueryRow(context.Background(), query, login)
	if err := row.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.AvatarId, &u.Description, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByEmail(conn *pgx.Conn, email string) error {
	query := `SELECT * FROM users WHERE email = $1`

	row := conn.QueryRow(context.Background(), query, email)
	if err := row.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.AvatarId, &u.Description, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser(conn *pgx.Conn, user *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) DeleteUser(conn *pgx.Conn, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) ForgetPassword(conn *pgx.Conn, password string) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`

	row := conn.QueryRow(context.Background(), query, password, &u.Email)
	if err := row.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.AvatarId, &u.Description, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (u *User) GetAvatar(conn *pgx.Conn, id int) error {
	query := `SELECT avatar_id FROM users WHERE id  =  $1`

	row := conn.QueryRow(context.Background(), query, id)
	if err := row.Scan(&u.AvatarId); err != nil {
		return err
	}

	return nil
}

func (u *User) GetSubscribers(conn *pgx.Conn, id int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetFollowers(conn *pgx.Conn, id int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUsers(conn *pgx.Conn) ([]User, error) {
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

func (u *User) UserSignIn(conn *pgx.Conn, email string, password string) error {
	query := `SELECT * FROM users WHERE email = $1 and password = $2`

	row := conn.QueryRow(context.Background(), query, email, password)
	if scanErr := row.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.AvatarId, &u.Description, &u.CreatedAt, &u.UpdatedAt); scanErr != nil {
		slog.Error("Error scanning row: %s", scanErr)
		return scanErr
	}

	return nil
}

func (u *User) UserSignUp(conn *pgx.Conn) error {
	query := `INSERT INTO users(login, email, password, created_at) VALUES($1, $2, $3, $4)`
	if _, err := conn.Exec(context.Background(), query, &u.Login, &u.Email, &u.Password, time.Now()); err != nil {
		slog.Error("Error creating user: %s", err)
		return err
	}

	return nil
}
