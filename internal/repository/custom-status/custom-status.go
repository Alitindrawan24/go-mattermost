package customstatus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/Alitindrawan24/go-mattermost/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func (r *Repository) UpdateCustomStatus(ctx *gin.Context, request entity.CustomStatusRequest) error {
	host := config.Get("MATTERMOST_HOST")
	accessToken := config.Get("MATTERMOST_ACCESS_TOKEN")
	userID := config.Get("MATTERMOST_USER_ID")

	url := host + "/api/v4/users/" + userID + "/status/custom"

	fmt.Println(url)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("PUT", url, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}

	if err != nil {
		fmt.Print(err.Error())
	}

	return nil
}
