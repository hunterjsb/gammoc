package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hunterjsb/gammoc/internal/client"
	"github.com/hunterjsb/gammoc/internal/config"
)

const (
	characterName  = "mubs" // Replace with your character's name
	mineX, mineY   = 2, 0
	smeltX, smeltY = 1, 5
	baseURL        = "https://api.artifactsmmo.com" // Replace with the actual base URL
)

type Client struct {
	http.Client
	Token string
}

type MovementRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ActionResponse struct {
	Data struct {
		Cooldown struct {
			TotalSeconds     int       `json:"totalSeconds"`
			RemainingSeconds int       `json:"remainingSeconds"`
			Expiration       time.Time `json:"expiration"`
		} `json:"cooldown"`
	} `json:"data"`
}

type ErrorResponse struct {
	Detail string `json:"detail"`
}

type CharacterInfo struct {
	Data struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"data"`
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := baseURL + endpoint
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}

func (c *Client) readErrorResponse(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Failed to read error response"
	}
	var errorResp ErrorResponse
	if err := json.Unmarshal(body, &errorResp); err != nil {
		return string(body)
	}
	return errorResp.Detail
}

func (c *Client) MoveCharacter(name string, x, y int) error {
	endpoint := fmt.Sprintf(string(client.EndpointActionMove), name)
	body := MovementRequest{X: x, Y: y}
	resp, err := c.makeRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorDetail := c.readErrorResponse(resp)
		return fmt.Errorf("unexpected status code: %d, detail: %s", resp.StatusCode, errorDetail)
	}

	var actionResp ActionResponse
	if err := json.NewDecoder(resp.Body).Decode(&actionResp); err != nil {
		return err
	}

	cooldown := actionResp.Data.Cooldown.RemainingSeconds
	fmt.Printf("Move cooldown: %d seconds\n", cooldown)
	time.Sleep(time.Duration(cooldown) * time.Second)

	return nil
}

func (c *Client) GatherResource(name string) error {
	endpoint := fmt.Sprintf(string(client.EndpointActionGathering), name)
	resp, err := c.makeRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorDetail := c.readErrorResponse(resp)
		return fmt.Errorf("unexpected status code: %d, detail: %s", resp.StatusCode, errorDetail)
	}

	var actionResp ActionResponse
	if err := json.NewDecoder(resp.Body).Decode(&actionResp); err != nil {
		return err
	}

	cooldown := actionResp.Data.Cooldown.RemainingSeconds
	fmt.Printf("Gathering cooldown: %d seconds\n", cooldown)
	time.Sleep(time.Duration(cooldown) * time.Second)

	return nil
}

func (c *Client) CraftItem(name string) error {
	endpoint := fmt.Sprintf(string(client.EndpointActionCrafting), name)
	resp, err := c.makeRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorDetail := c.readErrorResponse(resp)
		return fmt.Errorf("unexpected status code: %d, detail: %s", resp.StatusCode, errorDetail)
	}

	var actionResp ActionResponse
	if err := json.NewDecoder(resp.Body).Decode(&actionResp); err != nil {
		return err
	}

	cooldown := actionResp.Data.Cooldown.RemainingSeconds
	fmt.Printf("Crafting cooldown: %d seconds\n", cooldown)
	time.Sleep(time.Duration(cooldown) * time.Second)

	return nil
}

func (c *Client) GetCharacterInfo(name string) (CharacterInfo, error) {
	endpoint := fmt.Sprintf(string(client.EndpointGetCharacter), name)
	resp, err := c.makeRequest("GET", endpoint, nil)
	if err != nil {
		return CharacterInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorDetail := c.readErrorResponse(resp)
		return CharacterInfo{}, fmt.Errorf("unexpected status code: %d, detail: %s", resp.StatusCode, errorDetail)
	}

	var charInfo CharacterInfo
	if err := json.NewDecoder(resp.Body).Decode(&charInfo); err != nil {
		return CharacterInfo{}, err
	}

	return charInfo, nil
}

func main() {
	// Read the .env file
	env, err := config.ReadDotenv(".env")
	if err != nil {
		panic(err)
	}

	token := env["API_TOKEN"]
	if token == "" {
		panic("API_TOKEN not found in .env file")
	}

	c := NewClient(token)

	for {
		// Check current position
		charInfo, err := c.GetCharacterInfo(characterName)
		if err != nil {
			fmt.Printf("Error getting character info: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if charInfo.Data.X != mineX || charInfo.Data.Y != mineY {
			fmt.Println("Moving to mine")
			err := c.MoveCharacter(characterName, mineX, mineY)
			if err != nil {
				fmt.Printf("Error moving to mine: %v\n", err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		// Mine copper
		fmt.Println("Mining copper")
		err = c.GatherResource(characterName)
		if err != nil {
			fmt.Printf("Error mining copper: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Check current position again
		charInfo, err = c.GetCharacterInfo(characterName)
		if err != nil {
			fmt.Printf("Error getting character info: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if charInfo.Data.X != smeltX || charInfo.Data.Y != smeltY {
			fmt.Println("Moving to smelt")
			err = c.MoveCharacter(characterName, smeltX, smeltY)
			if err != nil {
				fmt.Printf("Error moving to smelt: %v\n", err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		// Smelt copper
		fmt.Println("Smelting copper")
		err = c.CraftItem(characterName)
		if err != nil {
			fmt.Printf("Error smelting copper: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Cycle completed, starting next cycle...")
		time.Sleep(5 * time.Second)
	}
}
