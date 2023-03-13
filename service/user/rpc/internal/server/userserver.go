// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/evpeople/ShopOCR/service/user/rpc/internal/logic"
	"github.com/evpeople/ShopOCR/service/user/rpc/internal/svc"
	"github.com/evpeople/ShopOCR/service/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
