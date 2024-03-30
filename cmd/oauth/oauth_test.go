package oauth

import (
	"testing"
)

func TestRequestAccessToken(t *testing.T) {
	err := GetConfiguration()
	if err != nil {
		t.Errorf("failed to get configuration for oauth; %v", err)
		t.FailNow()
	}

	client, err := RequestAccess()
	if err != nil {
		t.Errorf("failed to get access token: %v", err)
		t.FailNow()
	}

	_, err = client.CurrentUser()
	if err != nil {
		t.Errorf("failed to get current user: %v", err)
		t.FailNow()
	}
}
