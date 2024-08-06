package auth

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

var ContextKey = new(struct{})

// Read Account data from authorized request
func AccountOf(ctx context.Context) *object.Account {
	if cv := ctx.Value(ContextKey); cv == nil {
		return nil

	} else if account, ok := cv.(*object.Account); !ok {
		return nil

	} else {
		return account

	}
}
