package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Asliddin3/api-gateway/config"
	ub "github.com/Asliddin3/api-gateway/genproto/user"
)

type IServiceManager interface {
	UserService() ub.UserServiceClient
}

type ServiceManager struct {
	userService ub.UserServiceClient
}

func (s *ServiceManager) UserService() ub.UserServiceClient {
	return s.userService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &ServiceManager{
		userService: ub.NewUserServiceClient(connUser),
	}

	return serviceManager, nil
}
