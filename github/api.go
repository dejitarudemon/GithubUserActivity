package github

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (http.Response, error) {
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
