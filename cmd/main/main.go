package main

import (
	"fmt"
	"oauth"
	"os"
)

func main() {
	resp, err := oauth.RequestAccessToken()
	if err != nil {
		fmt.Printf("error requesting access token: %v", err)
		os.Exit(0)
	}

	accessToken := resp.AccessToken
	if accessToken == "" {
		fmt.Printf("access token is nil")
		os.Exit(0)
	}

	fmt.Printf("received access token: %s", accessToken)

	// Todo: when to use os.exit and when to use panic

}
