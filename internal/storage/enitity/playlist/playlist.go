package playlist

import (
	"context"
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
	UploadImage(conn *pgx.Conn, id int, imageId string) error
	GetPlaylistById(conn *pgx.Conn, id int) error
	UpdatePlaylist(conn *pgx.Conn, id int) error
	DeletePlaylist(conn *pgx.Conn, id int) error
}

var _ IPlaylist = (*Playlist)(nil)

func (p *Playlist) DeletePlaylist(conn *pgx.Conn, id int) error {
	query := "DELETE FROM playlists WHERE id = $1"
	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		slog.Error("Error deleting playlist: %s", err)
		return err
	}

	return nil
}

func (p *Playlist) GetPlaylists(conn *pgx.Conn) ([]Playlist, error) {
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
	query := `INSERT INTO playlists(title, description, image_id, created_at,updated_at, author_id, courses_ids)
			  VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := conn.Exec(context.Background(), query, &p.Title, &p.Description, &p.ImageId, time.Now(), time.Now(), &p.AuthorId, &p.CoursesIds)
	if err != nil {
		slog.Error("Error creating playlist: %s", err)
		return err
	}

	return nil
}

func (p *Playlist) UploadImage(conn *pgx.Conn, id int, imageId string) error {
	query := `UPDATE playlists SET image_id = $1 WHERE id = $2`

	_, err := conn.Exec(context.Background(), query, &imageId, &id)
	if err != nil {
		slog.Error("Error updating playlist: %s", err)
		return err
	}

	return nil
}

func (p *Playlist) GetPlaylistById(conn *pgx.Conn, id int) error {
	query := `SELECT  * FROM playlists WHERE id = $1`
	row := conn.QueryRow(context.Background(), query, &id)

	if err := row.Scan(&p.Id, &p.Title, &p.Description, &p.ImageId, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.CoursesIds); err != nil {
		slog.Error("Error scanning row: %s", err)
		return err
	}

	return nil
}

func (p *Playlist) UpdatePlaylist(conn *pgx.Conn, id int) error {
	query := `UPDATE playlists SET title  =  $1, description  =  $2, image_id  =  $3, updated_at  =  $4 WHERE id  =  $5`
	_, err := conn.Exec(context.Background(), query, &p.Title, &p.Description, &p.ImageId, time.Now(), &id)
	if err != nil {
		slog.Error("Error updating playlist: %s", err)
		return err
	}

	return nil
}
