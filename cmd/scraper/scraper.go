package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type Song struct {
	Artist string
	Name   string
}

type Scraper struct {
	Collector *colly.Collector
	URL       string
	Songs     []Song
}

func GetConfiguration() error {
	err := readConfiguration()
	if err != nil {
		return err
	}

	err = ValidateOAuthConfig()
	if err != nil {
		return err
	}

	return nil
}

func (scraper *Scraper) GetCollector() {
	collector := colly.NewCollector()

	scraper.Collector = collector
}

func (scraper *Scraper) GenerateURL(findType string, date string) {
	switch findType {
	case "hot_100":
		url := fmt.Sprintf("%s/%s", CFG.URL.Hot100, date)
		scraper.URL = url
		return
	case "billboard_200":
		url := fmt.Sprintf("%s/%s", CFG.URL.Billboard200, date)
		scraper.URL = url
		return
	case "billboard_200_global":
		url := fmt.Sprintf("%s/%s", CFG.URL.Billboard200Global, date)
		scraper.URL = url
		return
	case "billboard_japan_hot_100":
		url := fmt.Sprintf("%s/%s", CFG.URL.BillboardJapanHot100, date)
		scraper.URL = url
		return
	}

	log.Panicf("invalid chart type, received: %s", findType)
}

func (scraper *Scraper) GetSongs() error {
	scraper.Collector.OnHTML(".o-chart-results-list__item", func(element *colly.HTMLElement) {
		song := Song{}

		songName := element.ChildText("h3")
		artistName := element.ChildText(".c-label")
		if songName != "" {
			log.Infof("received song %s by artist %s", songName, artistName)
			song.Name = songName
			song.Artist = artistName
			scraper.Songs = append(scraper.Songs, song)
		}
	})

	err := scraper.Collector.Visit(scraper.URL)
	if err != nil {
		return err
	}

	return nil
}
