package service_integrations

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func UpdateReviewStatus(reviewID int64, status string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			err = errors.New("panic recovered")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseURL := "http://review-servise:8080/review/status/update"
	params := url.Values{}
	params.Add("review_id", strconv.FormatInt(reviewID, 10))
	params.Add("status", status)
	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, "PATCH", fullURL, nil)
	if err != nil {
		log.Printf("error creating request: %v", err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error executing request: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("review-service returned error: %d", resp.StatusCode)
		return errors.New("review-service returned status code: " + strconv.Itoa(resp.StatusCode))
	}

	log.Println("request to review-service succeeded")
	err = nil
	return nil
}
