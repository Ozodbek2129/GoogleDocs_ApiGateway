package service

import (
	"api_gateway/config"
	"api_gateway/genproto/docs"
	"api_gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UserServiceClient
	Docsservice() docs.DocsServiceClient
}

type serviceManagerImpl struct {
	userClient    user.UserServiceClient
	docsClient docs.DocsServiceClient
}

func (s *serviceManagerImpl) UserService() user.UserServiceClient {
	return s.userClient
}

func (s *serviceManagerImpl) Docsservice() docs.DocsServiceClient {
	return s.docsClient
}



func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connDocs, err := grpc.Dial(
		config.Load().DOCS_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		userClient:         user.NewUserServiceClient(connUser),
		docsClient:       docs.NewDocsServiceClient(connDocs),
	}, nil
}