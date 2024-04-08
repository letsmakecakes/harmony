package playlist

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"scraper"
)

type Playlist struct {
	ID          spotify.ID
	Name        string
	Client      *spotify.Client
	Description string
	Tracks      []spotify.ID
	TrackIDChan chan spotify.ID
}

type Song struct {
	Artist string
	Name   string
}

func (playlist *Playlist) CreatePlaylist(client *spotify.Client, name string, description string) error {
	user, err := client.CurrentUser()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	userID := user.ID

	emptyPlaylist, err := client.CreatePlaylistForUser(userID, name, description, false)
	if err != nil {
		return fmt.Errorf("failed to create playlist: %v", err)
	}

	playlist.ID = emptyPlaylist.ID
	playlist.Client = client
	playlist.Name = name
	playlist.Description = description

	return nil
}

func (playlist *Playlist) searchSong(query string) (spotify.ID, error) {
	searchResult, err := playlist.Client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		return "", fmt.Errorf("couldn't search the song: %v", err)
	}

	log.Infof("found %d track(s) for query %s", searchResult.Tracks.Total, query)

	if searchResult.Tracks.Total == 0 {
		return "", fmt.Errorf("no track found for query %s", query)
	}

	track := searchResult.Tracks.Tracks

	return track[0].ID, nil
}

func (playlist *Playlist) AddTracks() (string, error) {
	snapshotID, err := playlist.Client.AddTracksToPlaylist(playlist.ID, playlist.Tracks...)
	if err != nil {
		return "", fmt.Errorf("unable to add tracks to playlist")
	}

	return snapshotID, nil
}

func (playlist *Playlist) SearchSongs(songs ...scraper.Song) {
	for _, sng := range songs {
		query := fmt.Sprintf("artist:%s track:%s", sng.Artist, sng.Name)
		trackID, err := playlist.searchSong(query)
		if err != nil {
			log.Errorf("unable to search song: %v", err)
			continue
		}
		playlist.Tracks = append(playlist.Tracks, trackID)
	}
}
