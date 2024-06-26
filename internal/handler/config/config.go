package config

import (
	"CourseHub/internal/storage"
	"github.com/jackc/pgx/v5"
)

type CourseHub struct {
	ConnDB   *pgx.Conn
	Database Database
	Server   Server
}

type Server struct {
	Address      string `yaml:"address"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	Timeout      int    `yaml:"timeout"`
}

type Database struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ReadConfig(path string) (*CourseHub, error) {
	return nil, nil
}

// Создай метод который соберёт структуру с конфигурацией для работы с БД
func NewService(path string) (*CourseHub, error) {
	config, err := ReadConfig(path)
	if err != nil {
		return nil, err
	}

	conn, err := storage.NewConnection(config.Database.Address, config.Database.Username, config.Database.Password)
	if err != nil {
		return nil, err
	}

	return &CourseHub{
		ConnDB:   conn,
		Database: config.Database,
		Server:   config.Server,
	}, nil
}
