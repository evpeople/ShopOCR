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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterReply, err error) {

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	_, err = l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	if err != model.ErrNotFound {
		return nil, errorx.NewDefaultError("用户名已存在")
	}
	userinfo, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{Name: req.Username, Password: req.Password})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	id, err := userinfo.LastInsertId()
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, id)
	if err != nil {
		return nil, err
	}
	return &types.RegisterReply{
		Id:           id,
		Name:         req.Username,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
