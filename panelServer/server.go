package panelServer

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func CreateNewServer(port string , ){

	lis , err := net.Listen("tcp"  , port)

	if err !=nil{
		log.Fatalf("Error %v" , err)
	}

	s := Server {}

	grpcServer := grpc.NewServer()

	RegisterIndexServer()


}
