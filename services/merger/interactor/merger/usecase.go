package merger

import (
	"context"
)

type UseCase interface {
	MergeChunks(ctx context.Context, sha string) error
}
