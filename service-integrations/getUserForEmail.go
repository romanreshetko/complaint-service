package service_integrations

import (
	"complaint-service/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetUserForEmail(userID int64) (models.UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseURL := "http://auth-service:8080/user/email"
	params := url.Values{}
	params.Add("user_id", strconv.FormatInt(userID, 10))
	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		log.Printf("error creating request: %v", err)
		return models.UserInfo{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error executing request: %v", err)
		return models.UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("auth-service returned error: %d", resp.StatusCode)
		return models.UserInfo{}, errors.New("error getting user info")
	}

	var userInfo models.UserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Printf("error decoding user info: %v", err)
		return models.UserInfo{}, err
	}

	return userInfo, nil
}
