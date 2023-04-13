package uploader

import (
	"OMPFinex-CodeChallenge/pkg/errs"
	"OMPFinex-CodeChallenge/services/manager/entity"
	"bufio"
	"context"
	"io"
	"os"
	"time"
)

// GetImage Gives image base64
func (u *Uploader) GetImage(ctx context.Context, sha string) (io.Reader, error) {
	for i := 0; i < 5; i++ {
		//todo file manager must be moved to repo
		image, err := u.imageRepo.Get(ctx, sha)
		if err != nil {
			return nil, err
		}
		imageEnt := entity.ImageModelToEntity(*image)

		if imageEnt.Status != string(entity.ReadyStatus) {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		reader, err := u.readImage(ctx, imageEnt)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	return nil, errs.NewValidationError("image isn't ready")
}

func (u *Uploader) readImage(ctx context.Context, image entity.Image) (io.Reader, error) {
	file, err := os.Open(image.Data)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	// make a read buffer
	return bufio.NewReader(file), nil
}
