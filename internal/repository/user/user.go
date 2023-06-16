package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Alitindrawan24/go-mattermost/internal/entity"
	"github.com/Alitindrawan24/go-mattermost/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func (r *Repository) GetUser(ctx *gin.Context) (entity.User, error) {
	host := config.Get("MATTERMOST_HOST")
	accessToken := config.Get("MATTERMOST_ACCESS_TOKEN")

	url := host + "/api/v4/users/username/alt24"

	req, err := http.NewRequest("GET", url, nil)
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

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user entity.User
	json.Unmarshal(responseData, &user)

	return user, nil
}
