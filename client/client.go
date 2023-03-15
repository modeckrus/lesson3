package client

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	http.Client
	Addr string
}

type User struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Deleted *time.Time `json:"deleted"`
}

func (u *User) isExist() (bool, error) {
	if u.Deleted == nil {
		return true, nil
	}

	return false, nil
}

func (c *Client) Ping() (string, error) {
	req, err := http.NewRequest("GET", c.Addr+"/ping", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
