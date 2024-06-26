package playlist

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type Playlist struct {
	AuthorId    int
	CoursesIds  []int
	CreatedAt   time.Time
	Description string
	Id          int64
	ImageId     string
	Title       string
	UpdatedAt   time.Time
}

type CourseStatus string

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
