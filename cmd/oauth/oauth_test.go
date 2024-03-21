package oauth

import (
	"testing"
)

func TestRequestAccessToken(t *testing.T) {
	response, err := RequestAccessToken()

	if err != nil {
		t.Errorf("failed to get access token: %v", err)
	}

	if response.AccessToken == "" {
		t.Errorf("access token is empty")
	}

	if response.TokenType != "Bearer" {
		t.Errorf("invalid token type, expected Bearer, received: %s", response.TokenType)
	}

	if response.ExpiresIn < 0 {
		t.Errorf("invalid token expire period, received: %d", response.ExpiresIn)
	}
}
