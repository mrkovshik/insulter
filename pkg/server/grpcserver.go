package insulter

import (
	"github.com/mrkovshik/Insulter/grpc/proto"
	"context"
)

type GRPCServer struct{
	InsulterImplementation func(string) string
}


func (s *GRPCServer) Insult(ctx context.Context, req *proto.Value_IN) (*proto.Value_OUT, error) {
	return &proto.Value_OUT{
		Value_OUT: string(s.InsulterImplementation(proto.Value_IN)),
	}, nil
}
