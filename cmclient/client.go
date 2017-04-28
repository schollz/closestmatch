package cmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Connection is the BoltDB server instance
type Connection struct {
	Address string
}

// Open will load a connection to BoltDB
func Open(address string) (*Connection, error) {
	c := new(Connection)
	c.Address = address
	resp, err := http.Get(c.Address + "/uptime")
	if err != nil {
		return c, err
	}
	defer resp.Body.Close()
	return c, nil
}

func (c *Connection) Closest(searchString string) (match string, err error) {
	type QueryJSON struct {
		SearchString string `json:"s"`
	}

	payloadJSON := new(QueryJSON)
	payloadJSON.SearchString = searchString

	payloadBytes, err := json.Marshal(payloadJSON)
	if err != nil {
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/match", c.Address), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	type ResultJSON struct {
		Result string `json:"r"`
	}
	var r ResultJSON
	err = json.NewDecoder(resp.Body).Decode(&r)
	match = r.Result
	return
}

func (c *Connection) ClosestN(searchString string, n int) (matches []string, err error) {
	matches = []string{}
	type QueryJSON struct {
		SearchString string `json:"s"`
		N            int    `json:"n"`
	}

	payloadJSON := new(QueryJSON)
	payloadJSON.SearchString = searchString
	payloadJSON.N = n

	payloadBytes, err := json.Marshal(payloadJSON)
	if err != nil {
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/match", c.Address), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	type ResultJSON struct {
		Results []string `json:"r"`
	}
	var r ResultJSON
	err = json.NewDecoder(resp.Body).Decode(&r)
	matches = r.Results
	return
}
