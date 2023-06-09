package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/evpeople/ShopOCR/common/errorx"
	"github.com/evpeople/ShopOCR/service/search/api/internal/svc"
	"github.com/evpeople/ShopOCR/service/search/api/internal/types"
	"github.com/evpeople/ShopOCR/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchReply, error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, errorx.NewDefaultError("userId error")
	}

	// 使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}, nil
}
