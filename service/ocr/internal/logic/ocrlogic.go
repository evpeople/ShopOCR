package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	reader := NewImageReader()
	for {
		img, err := reader.ReadImg()
		if err != nil {
			break
		}
		callOCR("image", img)
		//TODO： 	修改架构，调用另一个python写的RPC，然后这个python的 PRC调用百度的OCR
	}
	return
}

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func callOCR(key, value string) {
	payload := Payload{key, value}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// 创建 POST 请求
	url := "https://example.com/api"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
