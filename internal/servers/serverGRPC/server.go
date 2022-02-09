/*
	Autor Andrey Semochkin
*/

package serverGRPC

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/deepch/Server/internal/constant"
	"github.com/deepch/Server/internal/memory"

	"google.golang.org/grpc/metadata"

	"github.com/deepch/Server/internal/function"
	"google.golang.org/grpc"

	pb "github.com/deepch/Server/api/proto/server"
)

/*

	service GRPC API

*/

//ServiceGRPC basic service structure
type ServiceGRPC struct {
	ServiceName        string
	Logger             *logrus.Logger
	WaitTaskFromSystem chan bool
	el                 *grpc.Server
}

//New - the function is executed when the service is initialized
func (service *ServiceGRPC) New() error {

	//If the server is disabled in the configuration
	if !memory.TnxMemory.DB.Module.API.GRPC.Enable {

		return nil

	}

	//TODO go to func
	//We will not start the server if the port is less than 100 and more than 65k
	if len(memory.TnxMemory.DB.Module.API.GRPC.Port) < 2 ||
		function.StringConvertToInt(memory.TnxMemory.DB.Module.API.GRPC.Port[1:]) < 100 ||
		function.StringConvertToInt(memory.TnxMemory.DB.Module.API.GRPC.Port[1:]) > 65535 {

		return constant.ErrorBadPortPort

	}

	//We hang the server on the specified port in the configuration
	l, err := net.Listen("tcp", memory.TnxMemory.DB.Module.API.GRPC.Port)

	if err != nil {

		return err

	}

	//Create a new grpc server
	s := grpc.NewServer()

	//Registering Methods
	pb.RegisterGreeterServer(s, &serverGRPC{service: service})

	go func() {

		err := s.Serve(l)

		if err != nil {
			service.Logger.Error(constant.ConstInitializationModuleGRPCError, err)
			service.WaitTaskFromSystem <- true
		}

	}()

	service.el = s

	return nil

}

//authentication - authorization check for GRPC server
func (service *ServiceGRPC) authentication(ctx context.Context) error {

	mb, _ := metadata.FromIncomingContext(ctx)

	authHeader := mb.Get(constant.ConstAuthorization)

	if len(authHeader) < 1 {

		return constant.ErrNoAuthHeader

	}

	if authHeader[0] != memory.TnxMemory.DB.Module.API.GRPC.Token {

		return constant.ErrIncorrectToken

	}

	return nil
}

//Name - returns the service name
func (service *ServiceGRPC) Name() string {

	return service.ServiceName

}

//Close - executed at the end of the service
func (service *ServiceGRPC) Close() error {

	if memory.TnxMemory.DB.Module.API.GRPC.Enable && service.el != nil {

		service.el.GracefulStop()

	}

	return nil

}
