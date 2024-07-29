package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hunterjsb/gammoc/internal/models"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func (c *Client) headers() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.Token,
	}
}

func (c *Client) GetCharacter(name string) (*models.Character, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("/characters/%s", name), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response models.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	character, ok := response.Data.(*models.Character)
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	return character, nil
}

func (c *Client) MoveCharacter(name string, x, y int) (*models.MovementResponse, error) {
	req := models.MovementRequest{X: x, Y: y}
	resp, err := c.makeRequest("POST", fmt.Sprintf("/my/%s/action/move", name), req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response models.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	movementResponse, ok := response.Data.(*models.MovementResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	return movementResponse, nil
}

func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := c.BaseURL + endpoint

	var req *http.Request

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	// Apply headers
	for key, value := range c.headers() {
		req.Header.Set(key, value)
	}

	return c.HTTPClient.Do(req)
}
