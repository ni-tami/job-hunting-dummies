package client

// // import vikstrous/dataloadgen with your other imports
// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/ni-tami/job-hunting-dummies-service/internal/model"
// 	"github.com/ni-tami/job-hunting-dummies-service/internal/repository"
// 	"github.com/vikstrous/dataloadgen"
// 	"gorm.io/gorm"
// )

// type ctxKey string

// const (
// 	loadersKey = ctxKey("dataloaders")
// )

// // Loaders wrap your data loaders to inject via middleware
// type Loaders struct {
// 	UserLoader        *dataloadgen.Loader[string, *model.User]
// 	ApplicationLoader *dataloadgen.Loader[string, *model.Application]
// }

// // NewLoaders instantiates data loaders for the middleware
// func NewLoaders(conn *gorm.DB) *Loaders {
// 	// define the data loader
// 	r := repository.NewJobPortalRepository(conn)
// 	return &Loaders{
// 		UserLoader:        dataloadgen.NewLoader(r.GetUsers, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
// 		ApplicationLoader: dataloadgen.NewLoader(r.GetApplicationsByApplicantID, dataloadgen.WithWait(time.Millisecond), dataloadgen.WithBatchCapacity(3)),
// 	}
// }

// // Middleware injects data loaders into the context
// func Middleware(conn *gorm.DB, next http.Handler) http.Handler {
// 	// return a middleware that injects the loader to the request context
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		loader := NewLoaders(conn)
// 		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
// 		next.ServeHTTP(w, r)
// 	})
// }

// // For returns the dataloader for a given context
// func For(ctx context.Context) *Loaders {
// 	return ctx.Value(loadersKey).(*Loaders)
// }

// // GetUser returns single user by id efficiently
// func GetUser(ctx context.Context, userID string) (*model.User, error) {
// 	loaders := For(ctx)
// 	return loaders.UserLoader.Load(ctx, userID)
// }

// // GetUsers returns many users by ids efficiently
// func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
// 	loaders := For(ctx)
// 	return loaders.UserLoader.LoadAll(ctx, userIDs)
// }

// // GetApplicationsByApplicantID returns many users by ids efficiently
// func GetApplicationsByApplicantID(ctx context.Context, applicantId int64) ([]*model.Application, error) {
// 	loaders := For(ctx)
// 	return loaders.ApplicationLoader.LoadAll(ctx, applicantId)
// }
