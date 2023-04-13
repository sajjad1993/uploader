package uploader

import (
	"context"
	"io"
)

type UseCase interface {
	GetImage(ctx context.Context, sha string) (io.Reader, error)
}
