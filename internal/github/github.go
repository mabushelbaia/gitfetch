package github

import (
	"encoding/json"
	"fmt"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"net/http"
	"os"
)

func LoadSampleUser() (*UserInfo, error) {
	data, err := os.ReadFile("sample_response.json")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info: %w", err)
	}

	user := UserInfo{}
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &user, nil
}

func FetchUserInfo(username string) (*UserInfo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info %w", err)

	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}
	user := UserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &user, nil
}
