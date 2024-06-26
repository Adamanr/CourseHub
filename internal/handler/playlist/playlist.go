package playlist

import (
	"CourseHub/internal/storage/enitity/playlist"
	playlistEntity "CourseHub/internal/storage/enitity/playlist"
	"github.com/jackc/pgx/v5"
)

type CourseHub struct {
	Conn *pgx.Conn
}

func GetPlaylists(conn *pgx.Conn) ([]playlist.Playlist, error) {
	playlists, err := playlist.GetPlaylists(conn)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (p *CourseHub) CreatePlaylist(playlist *playlistEntity.Playlist) error {
	if err := playlist.CreatePlaylist(p.Conn); err != nil {
		return err
	}

	return nil
}
