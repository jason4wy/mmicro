package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	hello "github.com/jason4wy/mmicro/hello/srv/proto/hello"
)

type Hello struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Hello) Call(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received Hello.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Hello) Stream(ctx context.Context, req *hello.StreamingRequest, stream hello.Hello_StreamStream) error {
	log.Logf("Received Hello.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&hello.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Hello) PingPong(ctx context.Context, stream hello.Hello_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&hello.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
