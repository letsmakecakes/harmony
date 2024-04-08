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
	//limit := 2

	scraper := new(Scraper)

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
}
