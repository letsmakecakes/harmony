package playlist

import (
	"scraper"
	"spotifyclient"
	"testing"
)

func TestPlaylist_CreatePlaylist(t *testing.T) {
	err := spotifyclient.GetConfiguration()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	spotifyClient, err := spotifyclient.RequestAccess()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	playlist := new(Playlist)

	err = playlist.CreatePlaylist(spotifyClient, "Billboard Hot 100", "Top billboard hot 100")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if playlist.ID == "" {
		t.Error("Playlist ID is empty")
		t.FailNow()
	}
}

func TestPlaylist_SearchSong(t *testing.T) {
	err := spotifyclient.GetConfiguration()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	spotifyClient, err := spotifyclient.RequestAccess()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	playlist := new(Playlist)
	playlist.Client = spotifyClient

	query := "artist:Bring Me The Horizon track:Kool Aid"

	songID, err := playlist.searchSong(query)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if songID == "" {
		t.Error("Song ID is empty")
		t.FailNow()
	}
}

func TestPlaylist_AddTracks(t *testing.T) {
	err := spotifyclient.GetConfiguration()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	spotifyClient, err := spotifyclient.RequestAccess()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	playlist := new(Playlist)

	err = playlist.CreatePlaylist(spotifyClient, "Billboard Hot 100", "Top billboard hot 100")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if playlist.ID == "" {
		t.Error("Playlist ID is empty")
		t.FailNow()
	}

	playlist.Tracks = append(playlist.Tracks, "11dFghVXANMlKmJXsNCbNl")

	playlist.AddTracks()
}

func TestPlaylist_BillboardPlaylist(t *testing.T) {
	err := spotifyclient.GetConfiguration()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = scraper.GetConfiguration()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	spotifyClient, err := spotifyclient.RequestAccess()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	playlist := new(Playlist)
	findType := "hot_100"
	date := ""
	//limit := 2

	scraper := new(scraper.Scraper)

	scraper.GetCollector()
	scraper.GenerateURL(findType, date)

	err = scraper.GetSongs()
	if err != nil {
		t.Errorf("failed to get songs list: %v", err)
		t.FailNow()
	}

	if len(scraper.Songs) == 0 {
		t.Errorf("songs list is empty")
	}

	err = playlist.CreatePlaylist(spotifyClient, "Billboard Hot 100", "Top billboard hot 100")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if playlist.ID == "" {
		t.Error("Playlist ID is empty")
		t.FailNow()
	}

	//playlist.Tracks = append(playlist.Tracks, "11dFghVXANMlKmJXsNCbNl")

	playlist.SearchSongs(scraper.Songs...)

	playlist.AddTracks()
}
