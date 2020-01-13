package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/jason4wy/mmicro/user/api/handler"
	"github.com/jason4wy/mmicro/user/api/client"

	user "github.com/jason4wy/mmicro/user/api/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.jason4wy.api.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the User srv client
		micro.WrapHandler(client.UserWrapper(service)),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
