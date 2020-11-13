package api

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const URL = "https://notify-api.line.me/api/notify"

func Notify(msg string) error {
	accessToken := os.Getenv("LINE_NOTIFY_TOKEN")
	if accessToken == "" {
		return errors.New("not found LINE_NOTIFY_TOKEN")
	}

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return err
}
