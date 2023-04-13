package donwloder

import (
	"OMPFinex-CodeChallenge/pkg/errs"
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
)

// RegisterImage register images to merges theirs chunks
func (m *Manager) RegisterImage(ctx context.Context, image entity.Image) error {
	image.Status = string(entity.UnCompletedStatus)
	imageModel := entity.ImageEntityToModel(image)
	//check duplicate image
	isDuplicate, err := m.imageRepo.CheckDuplicate(ctx, imageModel)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	if isDuplicate {
		//todo it must return 409 status
		return errs.NewConflictEntity("this image have already registered")
	}
	// save to repo
	err = m.imageRepo.Save(ctx, imageModel)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	return nil
}
