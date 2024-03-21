package artist

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"oauth"
)

type externalURLs struct {
	Spotify string `json:"spotify"`
}

type followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type image struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type Artist struct {
	ExternalURLs externalURLs `json:"external_urls"`
	Followers    followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []image      `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	TypeOfArtist string       `json:"type"`
	URI          string       `json:"uri"`
}

func GetArtist(artistID string, accessResponse oauth.AccessResponse) (Artist, error) {
	log.Infof("received artist id: %s", artistID)

	req, err := generateRequest(artistID, accessResponse)
	if err != nil {
		return Artist{}, err
	}

	log.Debugf("request generated: %v", req)

	resp, err := sendRequest(req)
	if err != nil {
		return Artist{}, err
	}

	log.Debugf("request has been sent, received response: %v", resp)

	decodedResponse, err := decodeResponse(resp)
	if err != nil {
		return Artist{}, err
	}

	log.Infof("decoded the response: %v", decodedResponse)

	return decodedResponse, nil
}

func generateRequest(id string, accessResponse oauth.AccessResponse) (*http.Request, error) {
	reqURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s", id)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return req, fmt.Errorf("error generating a request: %v", err)
	}

	authorizationValue := fmt.Sprintf("%s %s", accessResponse.TokenType, accessResponse.AccessToken)

	req.Header.Set("Authorization", authorizationValue)

	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}

func decodeResponse(resp *http.Response) (Artist, error) {
	var artist Artist

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return Artist{}, fmt.Errorf("could not parse JSON response: %v", err)
	}

	return artist, nil
}
