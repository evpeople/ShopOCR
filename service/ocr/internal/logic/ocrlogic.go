package logic

import (
	"context"

	"github.com/evpeople/ShopOCR/service/ocr/internal/svc"
	"github.com/evpeople/ShopOCR/service/ocr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OcrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOcrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OcrLogic {
	return &OcrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OcrLogic) Ocr(req *types.OcrReq) (resp *types.OcrReply, err error) {
	// todo: add your logic here and delete this line

	return
}
