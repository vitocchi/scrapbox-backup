package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

const ENCRYPTED_GITHUB_TOKEN = "AQICAHgOhA6XEO8DuSQDXosJ6s6ROlrTcFIxvwlBqAn6IjPbgQFNZpbEchSj8gi8G/4yQsLQAAAAhzCBhAYJKoZIhvcNAQcGoHcwdQIBADBwBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDJmLhaNTyHPfV8rdZwIBEIBDET5cmHx/8Up4Z/+R5mjwW/OEj4ja4TLPwQ0/uFVLyGbPFb2aHFjwg15lRv3pCGYTwsdhYMq6wQmQHD2f7DBIxVNM/Q=="

func handler() (int, error) {
	client, err := newScrapBoxClient()
	if err != nil {
		return 0, nil
	}
	return client.getLatestBackupID()
	/*
		scrapboxToken, err = getScrapboxToken()
		if err != nil {
			return error
		}
		cookie := &http.Cookie{Name: "connect.sid", Value: scrapboxToken}

		req, err := http.NewRequest("GET", "https://scrapbox.io/api/project-backup/vitocchi/list", nil)
		if err != nil {
			return nil, err
		}
		req.AddCookie(cookie)

		return req, nil

		githubToken, err := getGitHubToken()
		if err != nil {
			return error
		}
		return nil
	*/
}

func getGitHubToken() (string, error) {
	return decryptByKMS(ENCRYPTED_GITHUB_TOKEN)
}

func main() {
	lambda.Start(handler)
}
