package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ENCRYPTED_SCRAPBOX_AUTH = "AQICAHgOhA6XEO8DuSQDXosJ6s6ROlrTcFIxvwlBqAn6IjPbgQH6exap1IdKFfKx9CdEXbztAAAAtjCBswYJKoZIhvcNAQcGoIGlMIGiAgEAMIGcBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDBmIiddUuJUR44w8HAIBEIBvzGbTm5DwsM3ByQ/mGdAGng6gyfJ7IzAz7kQI+1OasqnqWP35epA7kXb0mrHjG7NcE3+tZSNwYNzUWDG1ieuNpqbhcaTPKuxvbSUwhgn0ymw2E+nEv7G7CU8twNVZBJtNmrOQTOnda1OjIBDols1H"

type scrapBoxClient struct {
	authCookie *http.Cookie
	client     *http.Client
}

func newScrapBoxClient() (*scrapBoxClient, error) {
	scrapboxToken, err := getScrapboxToken()
	if err != nil {
		return nil, err
	}
	return &scrapBoxClient{
		authCookie: &http.Cookie{Name: "connect.sid", Value: scrapboxToken},
		client:     &http.Client{},
	}, nil
}

func getScrapboxToken() (string, error) {
	return decryptByKMS(ENCRYPTED_SCRAPBOX_AUTH)
}

func (c *scrapBoxClient) getLatestBackupID() (int, error) {
	byteArray, err := c.throwGETRequest("https://scrapbox.io/api/project-backup/vitocchi/list")
	if err != nil {
		return 0, err
	}
	var res BackupResponse
	if err := json.Unmarshal(byteArray, &res); err != nil {
		return 0, err
	}
	return res.Backups[0].ID, nil
}

type BackupResponse struct {
	Backups      []Backup `json:"backups"`
	BackupEnable bool     `json:"backupEnable"`
}

type Backup struct {
	ID int `json:"backuped"`
}

func (c *scrapBoxClient) throwGETRequest(url string) ([]byte, error) {
	req, err := c.newGETRequest(url)
	if err != nil {
		return nil, err
	}
	return c.throwRequest(req)
}

func (c *scrapBoxClient) newGETRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(c.authCookie)
	return req, nil
}

func (c *scrapBoxClient) throwRequest(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
