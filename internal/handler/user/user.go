package user

import (
	userEntity "CourseHub/internal/storage/enitity/user"
	"github.com/jackc/pgx/v5"
)

func GetUser(conn *pgx.Conn) ([]userEntity.User, error) {
	users, err := userEntity.GetUsers(conn)
	if err != nil {
		return nil, err
	}

	return users, nil
}
