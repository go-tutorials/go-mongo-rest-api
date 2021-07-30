package app

import (
	"context"

	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func Route(r *mux.Router, context context.Context, root Root) error {
	app, err := NewApp(context, root)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)
	r.HandleFunc("/channel/{id}", app.TubeHandler.GetChannel).Methods(GET)
	r.HandleFunc("/channels/{id}", app.TubeHandler.GetChannels).Methods(GET)
	r.HandleFunc("/playlist/{id}", app.TubeHandler.GetPlaylist).Methods(GET)
	r.HandleFunc("/playlists/{id}", app.TubeHandler.GetPlaylists).Methods(GET)
	r.HandleFunc("/channelplaylists/{id}", app.TubeHandler.GetChannelPlaylists).Methods(GET)
	r.HandleFunc("/playlistvideos/{id}", app.TubeHandler.GetPlaylistVideos).Methods(GET)
	r.HandleFunc("/videos/{id}", app.TubeHandler.GetVideos).Methods(GET)

	return err
}
