package main

import (
	//"time"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/jason4wy/mmicro/hello/srv/handler"
	//"github.com/jason4wy/mmicro/hello/srv/subscriber"
	//"github.com/micro/go-plugins/registry/etcdv3"
	//"github.com/micro/go-micro/registry"

	hello "github.com/jason4wy/mmicro/hello/srv/proto/hello"
)

func main() {

/*	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
	       "http://127.0.0.1:2379",
	   }
	})
*/
	// New Service
	service := micro.NewService(
//		micro.Registry(reg),
		micro.Name("io.github.jason4wy.srv.hello"),
		micro.Version("latest"),
//                micro.RegisterTTL(time.Second*30),
//                micro.RegisterInterval(time.Second*15),

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
