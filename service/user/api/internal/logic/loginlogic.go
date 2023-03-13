package logic

import (
	"context"
	"strings"
	"time"

	"github.com/evpeople/ShopOCR/common/errorx"
	"github.com/evpeople/ShopOCR/service/user/api/internal/svc"
	"github.com/evpeople/ShopOCR/service/user/api/internal/types"
	"github.com/evpeople/ShopOCR/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginReply, error) {
	// check username and password
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	// find user by username
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}

	// check password
	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
