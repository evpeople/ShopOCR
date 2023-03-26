package logic

import (
	"bytes"
	"context"
	"encoding/base64"
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
	var encoded string
	if req.Url != "" {
		resp2, err := http.Get(req.Url)
		if err != nil {
			panic(err)
		}
		defer resp2.Body.Close()
		data, err := io.ReadAll(resp2.Body)
		if err != nil {
			panic(err)
		}
		encoded = base64.StdEncoding.EncodeToString(data)
	} else {
		encoded = req.Base64
	}
	c = make(chan string)
	// logx.Info(encoded)
	go callOCR("image", encoded)
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
	// var rep types.OcrReply
	// reply := `{"code":0,"msg":"","Replies":[{"content":"'The predicted text is :'","probability":0.9702153,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'上海斯格威铂尔-大酒店'","probability":0.9144976,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'0.870167'","probability":0.80206037,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'打浦路15号'","probability":0.88724154,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'0.959863'","probability":0.99456096,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'3'","probability":0.9988158,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'绿洲仕格维花园公寓'","probability":0.9155753,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'score:'","probability":0.9378295,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'订浦路252935号'","probability":0.9224631,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'score:'","probability":0.93276596,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'0.969589'","probability":0.99349326,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]},{"content":"'花费70.221455秒'","probability":0.8710235,"pos":[{"x":4,"y":6},{"x":186,"y":6},{"x":186,"y":18},{"x":4,"y":18}]}]}`
	// err = json.Unmarshal([]byte(reply), &rep)
	// if err != nil {
	// 	panic(err)
	// }
	// return &rep, nil
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
