package collector

import (
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
)

type UseCase interface {
	StartCollecting(ctx context.Context, image entity.Image)
	SaveChunk(ctx context.Context, chunk entity.Chunk) error
	CallMerger()
	GetChannel() chan string
}
