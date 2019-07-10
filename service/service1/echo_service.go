package echo

import (
	"context"
	"fmt"
)

type EchoService struct{}

func (s *EchoService) Echo(ctx context.Context, in *EchoIn) (out *EchoOut, err error) {
	fmt.Printf("in ------ %+v\n", in)
	return &EchoOut{Msg: "msg from server", Order: in.Order}, nil
}
