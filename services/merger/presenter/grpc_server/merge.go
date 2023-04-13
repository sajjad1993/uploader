package grpc_server

import (
	rpc "OMPFinex-CodeChallenge/pkg/rpc/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *MergerHandler) Merge(
	ctx context.Context,
	request *rpc.MergeRequest,
) (*rpc.MergeResponse, error) {
	err := h.interactor.MergeChunks(
		ctx,
		request.Image.Sha256,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &rpc.MergeResponse{Error: nil}, nil
}
