package main

import (
	"context"
	"errors"
	"grpc-demo/calculatorpb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(srv, &server{})
	reflection.Register(srv)

	serverErr := srv.Serve(listener)
	if serverErr != nil {
		panic(err)
	}

}

func (s *server) Calculate(ctx context.Context, request *calculatorpb.Request) (*calculatorpb.Response, error) {
	val1, val2 := request.GetValue1(), request.GetValue2()
	var result float32

	switch request.GetOperation() {
	case "add":
		result = val1 + val2
	case "sub":
		result = val1 - val2
	case "mult":
		result = val1 * val2
	case "div":
		if val2 == 0 {
			return nil, errors.New("division by zero not possible")
		}
		result = val1 / val2
	default:
		return nil, errors.New("operation not found")
	}

	return &calculatorpb.Response{Result: result}, nil
}
