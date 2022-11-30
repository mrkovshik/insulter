package insulter

import (
	"github.com/mrkovshik/insulter/grpc/proto"
	"context"
)

type GRPCServer struct{
	InsulterImplementation func() int
}

// Add method for return next Fibonacci
func (s *GRPCServer) Insult(ctx context.Context, req *proto.Value) (*proto.Value, error) {
	return &proto.Value{
		Value: int32(s.InsulterImplementation()),
	}, nil
}