package port

import "context"

type AuthorChecker interface {
	ExistsAll(ctx context.Context, ids []string) (bool, error)
}
