package handler

import (
	"api_gateway/genproto/doccs"
	"api_gateway/genproto/user"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.UserServiceClient
	Log         *slog.Logger
	DocsService doccs.DocsServiceClient
	Enforcer    *casbin.Enforcer
}

func NewHandler(user user.UserServiceClient, docs doccs.DocsServiceClient, logger *slog.Logger, Enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		UserService: user,
		Log:         logger,
		DocsService: docs,
		Enforcer: Enforcer,
	}
}
