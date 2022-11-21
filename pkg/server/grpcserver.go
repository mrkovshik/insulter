package calculus

import (
	"gitlab.com/tabby.ai/testing/tools/calculus/grpc/proto"
	"context"
)

type GRPCServer struct{
	FibonacciImplementation func() int
}

// Add method for return next Fibonacci
func (s *GRPCServer) Fibonacci(ctx context.Context, req *proto.Empty) (*proto.Value, error) {
	return &proto.Value{
		Value: int32(s.FibonacciImplementation()),
	}, nil
}