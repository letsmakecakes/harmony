package artist

import (
	"oauth"
	"testing"
)

func TestGetArtist(t *testing.T) {
	accessResponse, err := oauth.RequestAccessToken()
	if err != nil {
		t.Errorf("failed to get access token: %v", err)
		t.FailNow()
	}

	artistDetails, err := GetArtist("4Z8W4fKeB5YxbusRsdQVPb", accessResponse)

	if artistDetails.Name != "Radiohead" {
		t.Errorf("artist name does not match, expected Radiohead, received: %v", artistDetails)
	}
}
