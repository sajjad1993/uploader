package grpc_server

import (
	rpc "OMPFinex-CodeChallenge/pkg/rpc/proto"
	"OMPFinex-CodeChallenge/services/merger/interactor/merger"
)

type MergerHandler struct {
	rpc.UnimplementedMergerServer
	interactor merger.UseCase
}

func New(interactor merger.UseCase) *MergerHandler {
	return &MergerHandler{
		interactor: interactor,
	}
}
