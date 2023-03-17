package handler

import (
	"net/http"

	"github.com/evpeople/ShopOCR/service/ocr/api/internal/logic"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/svc"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ocrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OcrReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOcrLogic(r.Context(), svcCtx)
		resp, err := l.Ocr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
