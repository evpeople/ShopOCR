package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	reply := <-c
	fmt.Println(reply)
	type Reply struct {
		Err_no  int      `json:"err_no"`
		Err_msg string   `json:"err_msg"`
		Key     []string `json:"key"`
		Value   []string `json:"value"`
	}
	var r Reply
	err = json.Unmarshal([]byte(reply), &r)
	if err != nil {
		fmt.Println("Failed to unmarshal json:", err)
		return
	}
	// reply is
	//{"content":"{\"err_no\":0,\"err_msg\":\"\",\"key\":[\"result\"],\"value\":[\"[[('登机牌', 0.9866331), [[156.0, 27.0], [353.0, 24.0], [354.0, 67.0], [157.0, 70.0]]], [('BOARDING PASS', 0.921339), [[422.0, 23.0], [819.0, 15.0], [820.0, 55.0], [423.0, 63.0]]], [('序号SERIALNO.', 0.90068835), [[490.0, 103.0], [663.0, 101.0], [663.0, 120.0], [490.0, 122.0]]], [('CLASS', 0.91269666), [[398.0, 106.0], [455.0, 104.0], [456.0, 122.0], [399.0, 124.0]]], [('舱位', 0.9973184), [[343.0, 107.0], [385.0, 107.0], [385.0, 125.0], [343.0, 125.0]]], [('日期 DATE', 0.8522329), [[213.0, 108.0], [317.0, 107.0], [317.0, 127.0], [213.0, 128.0]]], [('座位号SEAT NO', 0.9227134), [[677.0, 99.0], [833.0, 96.0], [833.0, 116.0], [677.0, 119.0]]], [('航班 FLIGHT', 0.93869275), [[64.0, 112.0], [191.0, 108.0], [191.0, 128.0], [64.0, 132.0]]], [('W', 0.8183703), [[406.0, 132.0], [430.0, 132.0], [430.0, 157.0], [406.0, 157.0]]], [('035', 0.8816214), [[511.0, 130.0], [567.0, 130.0], [567.0, 155.0], [511.0, 155.0]]], [('O3DEC', 0.96435344), [[233.0, 138.0], [325.0, 136.0], [325.0, 157.0], [233.0, 159.0]]], [('MU2379', 0.99732256), [[83.0, 140.0], [212.0, 137.0], [212.0, 160.0], [83.0, 162.0]]], [('登机口', 0.8289368), [[489.0, 174.0], [553.0, 173.0], [553.0, 193.0], [490.0, 195.0]]], [('GATE', 0.99797285), [[566.0, 174.0], [612.0, 172.0], [613.0, 190.0], [567.0, 192.0]]], [('始发地', 0.99692935), [[343.0, 175.0], [409.0, 174.0], [410.0, 194.0], [344.0, 196.0]]], [('FROM', 0.9880327), [[404.0, 175.0], [468.0, 175.0], [468.0, 193.0], [404.0, 193.0]]], [('登机时间BDT', 0.9625783), [[678.0, 170.0], [810.0, 168.0], [810.0, 188.0], [678.0, 190.0]]], [('目的地TO', 0.9360944), [[67.0, 181.0], [168.0, 178.0], [168.0, 198.0], [68.0, 202.0]]], [('福州', 0.99901783), [[97.0, 207.0], [167.0, 206.0], [168.0, 227.0], [98.0, 229.0]]], [('TAIYUAN', 0.95021534), [[338.0, 219.0], [473.0, 216.0], [473.0, 235.0], [338.0, 239.0]]], [('G11', 0.6856511), [[505.0, 214.0], [553.0, 214.0], [553.0, 235.0], [505.0, 235.0]]], [('FUZHOU', 0.98853415), [[91.0, 231.0], [201.0, 227.0], [202.0, 248.0], [91.0, 251.0]]], [('身份识别ID NO', 0.8946389), [[345.0, 240.0], [482.0, 236.0], [482.0, 256.0], [345.0, 259.0]]], [('姓名NAME', 0.9741206), [[67.0, 251.0], [172.0, 249.0], [172.0, 268.0], [67.0, 270.0]]], [('ZHANGQIWET', 0.89279765), [[77.0, 278.0], [262.0, 274.0], [262.0, 294.0], [77.0, 297.0]]], [('票号 TKTNO', 0.9247335), [[462.0, 297.0], [578.0, 295.0], [578.0, 315.0], [462.0, 317.0]]], [('张祺伟', 0.9672672), [[103.0, 313.0], [208.0, 311.0], [208.0, 334.0], [103.0, 336.0]]], [('票价FARE', 0.9370944), [[70.0, 344.0], [164.0, 341.0], [165.0, 362.0], [70.0, 364.0]]], [('ETKT7813699238489/1', 0.9605225), [[346.0, 349.0], [660.0, 347.0], [660.0, 366.0], [346.0, 368.0]]], [('登机口于起飞前10分钟关闭', 0.91060454), [[102.0, 459.0], [344.0, 454.0], [344.0, 471.0], [102.0, 476.0]]], [('GATES CLOSE 1O MINUTES BEFORE DEPARTURE TIME', 0.91686845), [[357.0, 454.0], [830.0, 444.0], [830.0, 461.0], [357.0, 471.0]]]]\"],\"tensors\":[]}","probability":0.32,"pos":null}
	replies := buildReply(r.Value[0])
	return &types.OcrReply{Code: 0, Msg: "", Replies: replies}, nil
	//TODO： 	修改架构，调用另一个python写的RPC，然后这个python的 PRC调用百度的OCR
	// return &types.OcrReply{Content: r.Value[0], Confidence: 0.32, Pos: nil}, nil
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
