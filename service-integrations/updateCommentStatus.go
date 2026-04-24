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

func UpdateCommentStatus(commentID int64, status string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseURL := "http://comment-service:8080/comment/status/update"
	params := url.Values{}
	params.Add("comment_id", strconv.FormatInt(commentID, 10))
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
		log.Printf("comment-service returned error: %d", resp.StatusCode)
		return errors.New("comment returned status code: " + strconv.Itoa(resp.StatusCode))
	}

	log.Println("request to comment-service succeeded")
	return nil
}
