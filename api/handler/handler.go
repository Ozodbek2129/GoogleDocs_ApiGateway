package handler

import (
	"api_gateway/config"
	"api_gateway/genproto/user"
	"api_gateway/pkg/logger"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.UserServiceClient
	Log         *slog.Logger
	// Enforcer        *casbin.Enforcer
}

func NewHandler() *Handler {
	conf := config.Load()

	userr, err := grpc.Dial(conf.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	users := user.NewUserServiceClient(userr)

	return &Handler{
		UserService: users,
		Log:         logger.NewLogger(),
	}
}
