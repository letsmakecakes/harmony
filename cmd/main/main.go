package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"oauth"
	"os"
	"regexp"
	"scraper"
	"time"
)

// Todo: when to use os.exit and when to use panic

func realAllConfiguration() {
	err := scraper.GetConfiguration()
	if err != nil {
		log.Errorf("error getting configuration for scraping billboard songs: %v", err)
		os.Exit(0)
	}

	err = oauth.GetConfiguration()
	if err != nil {
		log.Errorf("failed to get configuration for oauth; %v", err)
		os.Exit(0)
	}
}

func isValidChartType(findType string) bool {
	return findType != "hot_100" || findType != "billboard_200" || findType != "billboard_200_global" || findType != "billboard_japan_hot_100"
}

func getAccessToken() string {
	resp, err := oauth.RequestAccessToken()
	if err != nil {
		log.Errorf("error requesting access token: %v", err)
		os.Exit(0)
	}

	return resp.AccessToken
}

func main() {
	realAllConfiguration()

	findType := flag.String("findType", "hot_100", "specify a chart type")
	currentDate := time.Now().Format("2006-01-02")
	date := flag.String("date", currentDate, "specify the chart date")
	songsLimit := flag.Int("limit", 1, "specify the number of songs")

	flag.Parse()

	if !isValidChartType(*findType) {
		log.Errorf("invalid chart type")
		os.Exit(0)
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(*date) {
		log.Errorf("invalid date format, please use yyyy-mm-dd")
		os.Exit(0)
	}

	if *songsLimit <= 0 {
		log.Errorf("invalid song limit, should greater than 0")
	}

	accessToken := getAccessToken()
	if accessToken == "" {
		log.Errorf("access token is nil")
		os.Exit(0)
	}

	log.Infof("received access token: %s", accessToken)

	url, err := scraper.GetURL(*findType, *date)
	if err != nil {
		log.Errorf("failed to get URL: %v", err)
		os.Exit(0)
	}

	collector := scraper.GetCollector()

	songs, err := scraper.GetSongs(collector, url)
	if err != nil {
		log.Errorf("failed to get/scrape songs: %v", err)
		os.Exit(0)
	}

	if len(songs) == 0 {
		log.Error("no songs received")
		os.Exit(0)
	}
}
