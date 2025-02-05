package grpc

import (
	"context"
	"example-svc/internal/some_domain/delivery"
	"example-svc/internal/some_domain/usecases/models"
	pb "example-svc/pkg/grpc"
)

type SomeGrpcHandler struct {
	pb.UnimplementedSomeExampleServer
	uc delivery.SomeDomain
}

func NewSomeGrpcHandler(uc delivery.SomeDomain) pb.SomeExampleServer {
	return &SomeGrpcHandler{
		uc: uc,
	}
}

func (h *SomeGrpcHandler) CreateSomething(ctx context.Context, req *pb.SomethingRequest) (*pb.SomethingResponse, error) {

	id, err := h.uc.SomeMethodCreate(ctx, models.CreateCommand{
		Param1: req.SomeField,
		Param2: "Filled",
	})
	if err != nil {
		return nil, err
	}

	return &pb.SomethingResponse{
		SomeField: id,
	}, nil
}
