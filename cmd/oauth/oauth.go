package oauth

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"net/http"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth       = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistModifyPrivate)
	clientChan = make(chan *spotify.Client)
	state      = "user"
)

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

func RequestAccess() (*spotify.Client, error) {
	startHTTPServer()

	auth.SetAuthInfo(CFG.Client.ID, CFG.Client.Secret)
	url := auth.AuthURL(state)

	log.Infof("please log in to spotify by visting the following page in your browser: %s", url)

	client := <-clientChan

	user, err := client.CurrentUser()
	if err != nil {
		return client, err
	}

	log.Infof("you are logged in as: %s", user.ID)

	return client, nil
}

func startHTTPServer() {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Infof("got request for: %s", request.URL.String())
	})
	go http.ListenAndServe(":8080", nil)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "couln't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("state mismatched: %s != %s\n", st, state)
	}

	client := auth.NewClient(token)

	fmt.Fprintf(w, "login completed!")

	clientChan <- &client
}
