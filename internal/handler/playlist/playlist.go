package playlist

import (
	playlistEntity "CourseHub/internal/storage/enitity/playlist"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type Hub struct {
	Conn *pgx.Conn
}

func (ch *Hub) DeletePlaylist(id int) error {
	playlist := playlistEntity.Playlist{}

	if err := playlist.DeletePlaylist(ch.Conn, id); err != nil {
		slog.Error("Error deleting playlist", "error", err)
		return err
	}

	return nil
}

type Handler interface {
	GetPlaylists() ([]playlistEntity.Playlist, error)
	CreatePlaylist(playlist playlistEntity.Playlist) (*playlistEntity.Playlist, error)
	UploadImage(id int, imageId string) (*playlistEntity.Playlist, error)
	GetPlaylistById(id int) (*playlistEntity.Playlist, error)
	UpdatePlaylist(playlist playlistEntity.Playlist, id int) (*playlistEntity.Playlist, error)
	DeletePlaylist(id int) error
}

var _ Handler = (*Hub)(nil)

func (ch *Hub) UpdatePlaylist(playlist playlistEntity.Playlist, id int) (*playlistEntity.Playlist, error) {
	if err := playlist.UpdatePlaylist(ch.Conn, id); err != nil {
		slog.Error("Error update playlist", err.Error)
		return nil, err
	}

	return &playlist, nil
}

func (ch *Hub) GetPlaylists() ([]playlistEntity.Playlist, error) {
	playlists := playlistEntity.Playlist{}

	return playlists.GetPlaylists(ch.Conn)
}

func (ch *Hub) CreatePlaylist(playlist playlistEntity.Playlist) (*playlistEntity.Playlist, error) {
	if err := playlist.CreatePlaylist(ch.Conn); err != nil {
		slog.Error("Error create playlist", err.Error)
		return nil, err
	}

	return &playlist, nil
}

func (ch *Hub) UploadImage(id int, imageId string) (*playlistEntity.Playlist, error) {
	playlist := playlistEntity.Playlist{}
	if err := playlist.UploadImage(ch.Conn, id, imageId); err != nil {
		slog.Error("Error upload image", err.Error)
		return nil, err
	}

	return &playlist, nil
}

func (ch *Hub) GetPlaylistById(id int) (*playlistEntity.Playlist, error) {
	playlist := playlistEntity.Playlist{}
	if err := playlist.GetPlaylistById(ch.Conn, id); err != nil {
		slog.Error("Error get playlist by id", err.Error)
		return nil, err
	}

	return &playlist, nil
}
