package playlist

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type Playlist struct {
	AuthorId    int       `json:"author_id"`
	CoursesIds  []int     `json:"courses_ids,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`
	Id          int64     `json:"id,omitempty"`
	ImageId     string    `json:"image_id,omitempty"`
	Title       string    `json:"title"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type IPlaylist interface {
	GetPlaylists(conn *pgx.Conn) ([]Playlist, error)
	CreatePlaylist(conn *pgx.Conn) error
}

func GetPlaylists(conn *pgx.Conn) ([]Playlist, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM playlists")
	if err != nil {
		slog.Error("Error querying users: %s", err)
		return nil, err
	}

	defer rows.Close()

	var playlist []Playlist
	for rows.Next() {
		var c Playlist

		if scanErr := rows.Scan(&c.Id, &c.Title, &c.Description, &c.ImageId, &c.CreatedAt, &c.UpdatedAt, &c.AuthorId, &c.CoursesIds); scanErr != nil {
			slog.Error("Error scanning row: %s", scanErr)
			return nil, scanErr
		}

		playlist = append(playlist, c)
	}

	return playlist, nil
}

func (p *Playlist) CreatePlaylist(conn *pgx.Conn) error {
	fmt.Println()
	fmt.Println(p.AuthorId)
	fmt.Println()

	query := `INSERT INTO playlists(title, description, image_id, created_at,updated_at, author_id, courses_ids)
			  VALUES($1, $2, $3, $4, $5, $6, $7)`

	rows, err := conn.Exec(context.Background(), query, &p.Title, &p.Description, &p.ImageId, time.Now(), time.Now(), &p.AuthorId, &p.CoursesIds)

	if err != nil {
		slog.Error("Error creating playlist: %s", err)
		return err
	}

	slog.Info("New playlist created", rows.String())

	return nil
}
