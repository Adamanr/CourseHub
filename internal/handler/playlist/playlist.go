package playlist

import (
	"CourseHub/internal/storage/enitity/playlist"
	"github.com/jackc/pgx/v5"
)

func GetPlaylists(conn *pgx.Conn) ([]playlist.Playlist, error) {
	playlists, err := playlist.GetPlaylists(conn)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}
