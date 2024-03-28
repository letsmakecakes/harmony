package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AccessResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func RequestAccessToken() (AccessResponse, error) {
	req, err := generateRequest()
	if err != nil {
		return AccessResponse{}, err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return AccessResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error while closing the response body: %v", err)
		}
	}(resp.Body)

	decodedResponse, err := getDecodedResponse(resp)
	if err != nil {
		return AccessResponse{}, err
	}

	return decodedResponse, nil
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

func generateRequest() (*http.Request, error) {
	body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", CFG.Client.ID, CFG.Client.Secret)
	requestBody := []byte(body)

	req, err := http.NewRequest("POST", CFG.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return req, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}

func getDecodedResponse(resp *http.Response) (AccessResponse, error) {
	var accessResponse AccessResponse

	if err := json.NewDecoder(resp.Body).Decode(&accessResponse); err != nil {
		return AccessResponse{}, fmt.Errorf("could not parse JSON response: %v", err)
	}

	return accessResponse, nil
}
