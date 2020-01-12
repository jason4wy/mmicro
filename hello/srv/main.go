package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"hello/srv/handler"
	"hello/srv/subscriber"

	hello "hello/srv/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("github.jason4wy.srv.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("github.jason4wy.srv.hello", service.Server(), new(subscriber.Hello))

	// Register Function as Subscriber
	micro.RegisterSubscriber("github.jason4wy.srv.hello", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
