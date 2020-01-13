package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/jason4wy/mmicro/user/srv/handler"
	//"github.com/jason4wy/mmicro/user/srv/subscriber"
	user "github.com/jason4wy/mmicro/user/srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.jason4wy.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("io.github.jason4wy.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("io.github.jason4wy.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
