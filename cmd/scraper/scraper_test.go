package scraper

import "testing"

func TestGetSongs(t *testing.T) {
	err := GetConfiguration()
	if err != nil {
		t.Errorf("failed to get configuartion: %v", err)
		t.FailNow()
	}

	findType := "hot_100"
	date := ""
	collector := GetCollector()

	url, err := GetURL(findType, date)
	if err != nil {
		t.Errorf("error getting url: %v", err)
		t.FailNow()
	}

	songs, err := GetSongs(collector, url)
	if err != nil {
		t.Errorf("failed to get songs list: %v", err)
		t.FailNow()
	}

	if len(songs) == 0 {
		t.Errorf("songs list is empty")
	}
}
