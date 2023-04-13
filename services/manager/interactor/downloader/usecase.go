package donwloder

import (
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
)

type UseCase interface {
	RegisterImage(ctx context.Context, image entity.Image) error
}
