package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

var c chan string

func NewOcrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OcrLogic {
	return &OcrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OcrLogic) Ocr(req *types.OcrReq) (resp *types.OcrReply, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("begin OCR:") // 这里的key和生成jwt token时传入的key一致
	reader := NewImageReader()
	img, err := reader.ReadImg()
	logx.Infof("read one img") // 这里的key和生成jwt token时传入的key一致
	logx.Info(img[:30])
	if err != nil {
		return nil, err
	}
	logx.Infof("call OCR") // 这里的key和生成jwt token时传入的key一致
	c = make(chan string)

	go callOCR("image", img)
	// logx.Info(<-c)

	//TODO： 	修改架构，调用另一个python写的RPC，然后这个python的 PRC调用百度的OCR
	return &types.OcrReply{Content: <-c, Confidence: 0.32, Pos: nil}, nil
}

type Payload struct {
	Key   []string `json:"key"`
	Value []string `json:"value"`
}

func callOCR(key, value string) {
	payload := Payload{[]string{key}, []string{value}}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// 创建 POST 请求
	url := "http://10.112.190.134:9998/ocr/prediction"
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
	// body, err := ioutil.ReadAll(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	logx.Infof("userId: %v", string(body)) // 这里的key和生成jwt token时传入的key一致

	c <- string(body)
}
