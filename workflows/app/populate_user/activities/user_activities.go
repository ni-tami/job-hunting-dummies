package activities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/models"
)

// FetchRandomUserActivity fetch randomuser.me api
func FetchRandomUserActivity() (models.CreateUser, error) {
	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		log.Println("No response from request")
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var randomUsers models.RandomUserResponse
	err = json.Unmarshal([]byte(body), &randomUsers)
	if err != nil {
		panic(err)
	}
	randUser := randomUsers.Results[len(randomUsers.Results)-1]
	user := randUser.ToGRPCUser()
	return user, nil
}
