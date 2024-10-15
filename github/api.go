package github

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const GITHUBURL = "https://api.github.com/"
const GITHUBEVENTS = GITHUBURL + "users/"
const EVENTS = "/events?page="
const ENDCONTENT = "5"
const LENGTH = "conten-length"

func get(url string) (http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		strings.NewReader(""),
	)
	defer cancel()

	if err != nil {
		return http.Response{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil
}

func GetEvenets(username string) (Events, error) {
	if len(username) == 0 {
		return Events{}, errors.New("USERNAME IS MISSED")
	}

	events := make(Events, 0)
	for i := 1; ; i++ {
		url := GITHUBEVENTS + username + EVENTS + strconv.Itoa(i)
		resp, err := get(url)
		if err != nil {
			return Events{}, err
		}

		if resp.StatusCode != 200 {
			return Events{}, errors.New("CAN'T CONNECT TO GITHUB")
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return Events{}, err
		}

		eventsOnPage := new(Events)
		err = json.Unmarshal(data, eventsOnPage)
		if err != nil {
			return Events{}, err
		}

		if len(*eventsOnPage) == 0 {
			return events, nil
		}

		events = append(events, *eventsOnPage...)
	}

}
