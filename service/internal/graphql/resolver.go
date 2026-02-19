// nolint:gci
package graphql

import "github.com/ni-tami/service/internal/graphql/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	todos []*model.Todo
}
