package insulter

import (
	"github.com/mrkovshik/insulter/grpc/proto"
	"context"
)

type GRPCServer struct{
	}


func (s *GRPCServer) Insult(ctx context.Context, req *proto.Value_IN) (*proto.Value_OUT, error) {
	var x proto.Value_OUT
	x.Value=insult(req.Value)
	return  &x, nil
}
