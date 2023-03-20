package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/evpeople/ShopOCR/common/errorx"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/svc"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OcrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var c chan string

func NewOcrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OcrLogic {
	return &OcrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OcrLogic) Ocr(req *types.OcrReq) (resp *types.OcrReply, err error) {
	c = make(chan string)

	go callOCR("image", req.Base64)
	reply := <-c
	type Reply struct {
		Err_no  int      `json:"err_no"`
		Err_msg string   `json:"err_msg"`
		Key     []string `json:"key"`
		Value   []string `json:"value"`
	}
	var r Reply
	err = json.Unmarshal([]byte(reply), &r)
	if err != nil {
		errorx.NewDefaultError(err.Error())
		return
	}
	replies := buildReply(r.Value[0])
	return &types.OcrReply{Code: 0, Msg: "", Replies: replies}, nil
}

type Payload struct {
	Key   []string `json:"key"`
	Value []string `json:"value"`
}

func callOCR(key, value string) {
	payload := Payload{[]string{key}, []string{value}}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}

	// 创建 POST 请求
	url := "http://10.112.190.134:9998/ocr/prediction"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
//	logx.Infof("userId: %v", string(body))

	c <- string(body)
}
