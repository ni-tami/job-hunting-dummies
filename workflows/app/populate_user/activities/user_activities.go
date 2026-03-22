package activities

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	pb "github.com/ni-tami/job-hunting-dummies-workflows/app/pb/out/go/job_hunting_dummies"
	"github.com/ni-tami/job-hunting-dummies-workflows/app/populate_user/models"
)

type JobHuntServiceActivity struct {
	grpcClient pb.JobHuntServiceClient
}

func NewJobHuntServiceActivity(grpcClient pb.JobHuntServiceClient) *JobHuntServiceActivity {
	return &JobHuntServiceActivity{
		grpcClient: grpcClient,
	}
}

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

func (a *JobHuntServiceActivity) CreateUserRPCActivity(ctx context.Context, randUser models.CreateUser) error {
	// Contact the server and print out its response.
	u, err := a.grpcClient.CreateUser(ctx, &pb.CreateUserRequest{Name: randUser.Name, Username: randUser.Username})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		return err
	}
	log.Printf("New User: %+v", u)
	return nil
}
