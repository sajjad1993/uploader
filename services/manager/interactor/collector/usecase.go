package interactor

import (
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
)

type UseCase interface {
	RegisterImage(ctx context.Context, image entity.Image) error
	SaveChunk(ctx context.Context, chunk entity.Chunk) error
	GetImage()
}
