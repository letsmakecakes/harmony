package scraper

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type Song struct {
	Artist string
	Name   string
}

var Songs []Song

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

func GetCollector() *colly.Collector {
	collector := colly.NewCollector()

	return collector
}

func GetURL(findType string, date string) (string, error) {
	switch findType {
	case "hot_100":
		url := fmt.Sprintf("%s/%s", CFG.URL.Hot100, date)
		return url, nil
	case "billboard_200":
		url := fmt.Sprintf("%s/%s", CFG.URL.Billboard200, date)
		return url, nil
	case "billboard_200_global":
		url := fmt.Sprintf("%s/%s", CFG.URL.Billboard200Global, date)
		return url, nil
	case "billboard_japan_hot_100":
		url := fmt.Sprintf("%s/%s", CFG.URL.BillboardJapanHot100, date)
		return url, nil
	}

	return "", errors.New("invalid find")
}

func GetSongs(collector *colly.Collector, url string) ([]Song, error) {
	collector.OnHTML(".o-chart-results-list__item", func(element *colly.HTMLElement) {
		song := Song{}

		songName := element.ChildText("h3")
		artistName := element.ChildText(".c-label")
		if songName != "" {
			log.Infof("received songName %s by artist %s", songName, artistName)
			song.Name = songName
			song.Artist = artistName
			Songs = append(Songs, song)
		}
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return Songs, nil
}
