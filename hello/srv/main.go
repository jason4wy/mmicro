package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/jason4wy/mmicro/hello/srv/handler"
	//"github.com/jason4wy/mmicro/hello/srv/subscriber"

	hello "github.com/jason4wy/mmicro/hello/srv/proto/hello"
)

func main() {

	// New Service
	service := micro.NewService(
		micro.Name("io.github.jason4wy.srv.hello"),
		micro.Version("latest"),

	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("io.github.jason4wy.srv.hello", service.Server(), new(subscriber.Hello))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("io.github.jason4wy.srv.hello", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
