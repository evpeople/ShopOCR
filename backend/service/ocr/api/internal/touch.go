package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/evpeople/ShopOCR/common/errorx"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PairData [4][2]int

// 获取Paris配对
func findMatchingBrackets(data string, symbolBegin, symbolEnd rune) ([][]int, error) {

	var stack []int
	var pairs [][]int
	// logx.Info(data)
	for i, char := range data {
		if char == symbolBegin {
			stack = append(stack, i)
		} else if char == symbolEnd {
			if len(stack) == 0 {
				return nil, errorx.NewDefaultError("unmatched bracket")
			}
			openIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairs = append(pairs, []int{openIndex, i})
		}
	}

	if len(stack) > 0 {
		return nil, errorx.NewDefaultError("unmatched bracket")
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	return pairs, nil
}
func buildReply(data string) (replies []types.SingleOcrReply) {
	// pairs1 和 pairs2 分别是两个括号的配对的序列
	// logx.Info(data)
	pairs2, err := findMatchingBrackets(data, '[', ']')
	logx.Info("pairs 2", pairs2)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
	pairs, err := findMatchingBrackets(data, '(', ')')
	logx.Info("pairs 1", pairs)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
	// 构造识别出的内容和可信度
	contents, confidences, err := buildMsg(data, pairs)
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
	//	构造识别出的位置
	poses := buildPos(data, pairs2)
	for i := 0; i < len(confidences); i++ {
		replies = append(replies, types.SingleOcrReply{Confidence: confidences[i], Content: contents[i], Pos: poses[i]})
	}
	return replies
}

func buildMsg(data string, pairs [][]int) (content []string, confidence []float64, err error) {
	for i := 0; i < len(pairs); i++ {
		originValue := fmt.Sprint(data[pairs[i][0]+1 : pairs[i][1]])
		values := strings.Split(originValue, ",")
		content = append(content, values[0])
		values[1] = strings.Trim(values[1], " ")
		tmpConfidence, err2 := strconv.ParseFloat(values[1], 64)
		confidence = append(confidence, tmpConfidence)
		if err2 != nil {
			errorx.NewDefaultError(err2.Error())
			err = err2
			return
		}
	}
	return
}
func buildPos(data string, pairs [][]int) [][]types.Position {
	pairs = pairs[1:]
	pS := len(pairs) / 6
	AllPairData := make([]PairData, pS)
	for i := 0; i < pS; i++ {
		AllPairData[i] = PairData{
			[2]int{pairs[i*6+2][0], pairs[i*6+2][1]}, [2]int{pairs[i*6+3][0], pairs[i*6+3][1]},
			[2]int{pairs[i*6+4][0], pairs[i*6+4][1]}, [2]int{pairs[i*6+5][0], pairs[i*6+5][1]}}
	}
	var ans [][]types.Position
	for j := 0; j < pS; j++ {
		var poses []types.Position
		for i := 0; i < 4; i++ {
			var pos types.Position
			values := fmt.Sprint(data[AllPairData[j][i][0]+1 : AllPairData[j][i][1]])
			xy := strings.Split(values, ",")
			x, err := strconv.ParseFloat(xy[0], 64)
			if err != nil {
				errorx.NewDefaultError(err.Error())
			}
			xy[1] = strings.Trim(xy[1], " ")
			y, err := strconv.ParseFloat(xy[1], 64)
			if err != nil {
				errorx.NewDefaultError(err.Error())
			}
			pos.X = x
			pos.Y = y
			poses = append(poses, pos)
		}
		ans = append(ans, poses)
	}
	return ans
}
func main() {
	data := `	[[('RunninganalysisFinference_op_replace_pass', 0.9174677), [[36.0, 3.0], [384.0, 3.0], [384.0, 15.0], [36.0, 15.0]]], [('Runninganalysis[memory_optimize_pass]', 0.96022177), [[31.0, 19.0], [347.0, 21.0], [347.0, 36.0], [31.0, 35.0]]], [('I0308 09:47:57.704764 65137 memory_optimize_pass.cc:200J Cluster name', 0.9541079), [[4.0, 39.0], [560.0, 39.0], [560.0, 55.0], [4.0, 55.0]]], [(':conv2d 89.tmp0size:153600', 0.9287803), [[558.0, 40.0], [810.0, 40.0], [810.0, 51.0], [558.0, 51.0]]], [('I030809:47:57.70478265137memory_optimize_pass.cc:200j', 0.9778392), [[4.0, 58.0], [449.0, 58.0], [449.0, 74.0], [4.0, 74.0]]], [('Cluster name:', 0.8931999), [[458.0, 59.0], [582.0, 59.0], [582.0, 72.0], [458.0, 72.0]]], [('elementwise_add_7size:358400', 0.97442114), [[577.0, 59.0], [827.0, 59.0], [827.0, 72.0], [577.0, 72.0]]], [('I030809:47:57.70478765137memory_optimize_pass.cc:200j', 0.82393646), [[3.0, 77.0], [450.0, 77.0], [450.0, 93.0], [3.0, 93.0]]], [('Cluster name', 0.9409812), [[456.0, 78.0], [557.0, 78.0], [557.0, 91.0], [456.0, 91.0]]], [('conv2d_90.tmp_0size:614400', 0.9291638), [[578.0, 78.0], [811.0, 78.0], [811.0, 91.0], [578.0, 91.0]]], [('I030809:47:57.70479265137memory_optimize_pass.cc:200j', 0.84824795), [[3.0, 95.0], [450.0, 96.0], [450.0, 112.0], [3.0, 111.0]]], [('Cluster name', 0.86440545), [[457.0, 95.0], [559.0, 98.0], [559.0, 112.0], [457.0, 110.0]]], [('batch_norm_48.tmp_2size:9830400', 0.9431684), [[576.0, 96.0], [851.0, 95.0], [851.0, 110.0], [576.0, 111.0]]], [('I030809:47:57.70479565137memory_optimize_pass.cc:200j', 0.833032), [[3.0, 115.0], [450.0, 115.0], [450.0, 131.0], [3.0, 131.0]]], [('Cluster name', 0.89548486), [[456.0, 115.0], [558.0, 116.0], [558.0, 131.0], [456.0, 129.0]]], [(':relu_2.tmp_0size:13107200', 0.94630444), [[556.0, 115.0], [804.0, 115.0], [804.0, 128.0], [556.0, 129.0]]], [('I030809:47:57.70479965137memory_optimize_pass.cc:200j', 0.83385575), [[3.0, 134.0], [450.0, 134.0], [450.0, 150.0], [3.0, 150.0]]], [('Cluster name', 0.9535084), [[457.0, 134.0], [558.0, 134.0], [558.0, 150.0], [457.0, 150.0]]], [('conv2d_96.tmp_0size:2457600', 0.933542), [[579.0, 135.0], [819.0, 135.0], [819.0, 148.0], [579.0, 148.0]]], [('I030809:47:57.70480365137memory_optimize_pass.cc:200j', 0.83004147), [[3.0, 153.0], [450.0, 153.0], [450.0, 169.0], [3.0, 169.0]]], [('Cluster', 0.889619), [[457.0, 154.0], [520.0, 154.0], [520.0, 167.0], [457.0, 167.0]]], [('name', 0.88112724), [[514.0, 155.0], [557.0, 155.0], [557.0, 167.0], [514.0, 167.0]]], [('conv2d_92.tmp_0size:9830400', 0.9343155), [[578.0, 153.0], [819.0, 152.0], [819.0, 167.0], [578.0, 167.0]]], [('I030809:47:57.70480765137memory_optimize_pass.cc:200', 0.94844234), [[2.0, 171.0], [450.0, 171.0], [450.0, 188.0], [2.0, 187.0]]], [('Cluster', 0.9368461), [[457.0, 171.0], [518.0, 173.0], [518.0, 187.0], [457.0, 186.0]]], [('name', 0.9913155), [[521.0, 175.0], [557.0, 175.0], [557.0, 187.0], [521.0, 187.0]]], [('tmp_1size:2457600', 0.8409946), [[579.0, 172.0], [741.0, 172.0], [741.0, 188.0], [579.0, 188.0]]], [('I030809:47:57.70481165137memory_optimize_pass.cc:200J', 0.8399474), [[3.0, 190.0], [451.0, 191.0], [451.0, 207.0], [3.0, 206.0]]], [('Cluster', 0.9011577), [[453.0, 192.0], [518.0, 192.0], [518.0, 206.0], [453.0, 206.0]]], [('name', 0.9032968), [[520.0, 194.0], [557.0, 194.0], [557.0, 205.0], [520.0, 205.0]]], [('size:4915200', 0.93415093), [[587.0, 192.0], [708.0, 191.0], [708.0, 204.0], [587.0, 206.0]]], [('---Running analysis[ir_graph_to_program_pass]', 0.96679944), [[4.0, 210.0], [378.0, 210.0], [378.0, 226.0], [4.0, 226.0]]], [('I0308 09:47:57.706780 65160 memory_optimize_pass.cc:200JCluster name:conv2d_89.tmp_0size:153600', 0.9609726), [[4.0, 229.0], [812.0, 229.0], [812.0, 245.0], [4.0, 245.0]]], [('I030809:47:57.70680165160memory_optimize_pass.cc:200JClustername', 0.9414423), [[3.0, 247.0], [563.0, 247.0], [563.0, 265.0], [3.0, 265.0]]], [('elementwise_add_7size:358400', 0.95025694), [[578.0, 248.0], [827.0, 248.0], [827.0, 262.0], [578.0, 262.0]]], [('I030809:47:57.70680765160memory_optimize_pass.cc:200JClustername', 0.92734325), [[3.0, 266.0], [561.0, 267.0], [561.0, 283.0], [3.0, 283.0]]], [('conv2d_90.tmp_0size:614400', 0.9257875), [[579.0, 267.0], [811.0, 267.0], [811.0, 281.0], [579.0, 281.0]]], [('I030809:47:57.70681365160memory_optimize_pass.cc:200]Clustername', 0.9628159), [[1.0, 283.0], [563.0, 284.0], [563.0, 303.0], [1.0, 303.0]]], [('batch_norm_48.tmp_2size:9830400', 0.92582285), [[577.0, 287.0], [850.0, 287.0], [850.0, 300.0], [577.0, 300.0]]], [('I030809:47:57.70681765160memory_optimize_pass.cc:200]Clustername', 0.92927533), [[2.0, 303.0], [559.0, 304.0], [559.0, 321.0], [2.0, 320.0]]], [('relu_2.tmp_0size:13107200', 0.9612367), [[577.0, 306.0], [804.0, 306.0], [804.0, 319.0], [577.0, 319.0]]], [('I030809:47:57.70682565160memory_optimize_pass.cc:200JCluster', 0.9252313), [[2.0, 322.0], [521.0, 323.0], [521.0, 340.0], [2.0, 339.0]]], [('name', 0.8784002), [[514.0, 326.0], [557.0, 326.0], [557.0, 338.0], [514.0, 338.0]]], [('conv2d_96.tmp_0size:2457600', 0.9434261), [[577.0, 323.0], [819.0, 322.0], [819.0, 338.0], [577.0, 339.0]]], [('I030809:47:57.70683165160memory_optimize_pass.cc:200J', 0.9292345), [[3.0, 342.0], [453.0, 343.0], [453.0, 358.0], [3.0, 357.0]]], [('cluster', 0.8343373), [[446.0, 342.0], [520.0, 344.0], [520.0, 358.0], [446.0, 356.0]]], [('name', 0.979149), [[514.0, 346.0], [557.0, 346.0], [557.0, 357.0], [514.0, 357.0]]], [('conv2d_92.tmp_0size:9830400', 0.903206), [[578.0, 343.0], [818.0, 343.0], [818.0, 357.0], [578.0, 357.0]]], [('I030809:47:57.70683665160memory_optimize_pass.cc:200J', 0.97951746), [[4.0, 362.0], [450.0, 362.0], [450.0, 378.0], [4.0, 378.0]]], [('CLuster', 0.8492581), [[452.0, 361.0], [518.0, 362.0], [518.0, 377.0], [452.0, 375.0]]], [('name', 0.8857137), [[520.0, 364.0], [557.0, 364.0], [557.0, 376.0], [520.0, 376.0]]], [('tmp_1size:2457600', 0.9388232), [[577.0, 362.0], [739.0, 362.0], [739.0, 376.0], [577.0, 377.0]]], [('I030809:47:57.70684165160memory_optimize_pass.cc:200J', 0.9177514), [[3.0, 379.0], [451.0, 380.0], [451.0, 397.0], [3.0, 396.0]]], [('Cluster name:xsize:4915200', 0.9554612), [[452.0, 381.0], [708.0, 381.0], [708.0, 397.0], [452.0, 397.0]]], [('I0308 09:47:57.708473 65173 analysis_predictor.cc:548', 0.9833433), [[4.0, 398.0], [433.0, 399.0], [433.0, 415.0], [4.0, 414.0]]], [('optimizeend', 0.8845839), [[504.0, 402.0], [612.0, 402.0], [612.0, 414.0], [504.0, 414.0]]], [('I030809:47:57.70853465173naive_executor.cc:107]-=-', 0.8840975), [[4.0, 418.0], [436.0, 420.0], [436.0, 434.0], [4.0, 432.0]]], [('skip [feed], feed -\u003e x', 0.9160608), [[450.0, 419.0], [631.0, 419.0], [631.0, 435.0], [450.0, 435.0]]], [('---Runninganalysis[ir_graph_to_program_pass]', 0.91975534), [[4.0, 439.0], [377.0, 439.0], [377.0, 453.0], [4.0, 453.0]]], [('I030809:47:57.71059265173naive_executor.cc:107', 0.92931813), [[5.0, 458.0], [407.0, 458.0], [407.0, 471.0], [5.0, 471.0]]], [('skip [save_infer_model/scale_0.tmp_0], fetch -\u003e fetch', 0.95083696), [[450.0, 457.0], [876.0, 457.0], [876.0, 473.0], [450.0, 473.0]]], [('I0308 09:47:57.743366 65137 analysis_predictor.cc:548', 0.9185199), [[4.0, 475.0], [432.0, 476.0], [432.0, 490.0], [4.0, 490.0]]], [('optimize end=', 0.93721956), [[505.0, 478.0], [629.0, 478.0], [629.0, 490.0], [505.0, 490.0]]], [('I030809:47:57.74344365137naive_executor.cc:107]-.', 0.8761739), [[4.0, 494.0], [425.0, 496.0], [425.0, 510.0], [4.0, 508.0]]], [('skip [feed], feed -\u003e x', 0.91317075), [[450.0, 495.0], [632.0, 495.0], [632.0, 511.0], [450.0, 511.0]]], [('I0308 09:47:57.74539565137 naive_executor.cc:107]', 0.9601425), [[4.0, 514.0], [406.0, 514.0], [406.0, 530.0], [4.0, 530.0]]], [('skip [relu_2.tmp_0], fetch -\u003e fetch', 0.8939199), [[451.0, 515.0], [731.0, 515.0], [731.0, 529.0], [451.0, 529.0]]], [('I0308 09:47:57.752557 65160 analysis_predictor.cc:548', 0.92047906), [[3.0, 532.0], [432.0, 534.0], [432.0, 549.0], [3.0, 547.0]]], [('optimizeend=', 0.8884151), [[504.0, 535.0], [634.0, 535.0], [634.0, 546.0], [504.0, 546.0]]], [('I030809:47:57.75262465160naive_executor.cc:107', 0.9371465), [[4.0, 551.0], [402.0, 553.0], [402.0, 566.0], [4.0, 565.0]]], [('skip [feed],feed -\u003e x', 0.91677195), [[450.0, 551.0], [630.0, 552.0], [630.0, 568.0], [450.0, 567.0]]], [('I030809:47:57.75458265160naive_executor.cc:107', 0.9265798), [[4.0, 570.0], [401.0, 571.0], [401.0, 586.0], [4.0, 585.0]]], [('skip [relu_2.tmp_o], fetch -\u003e fetch', 0.94723773), [[450.0, 571.0], [732.0, 571.0], [732.0, 587.0], [450.0, 587.0]]]]`
	buildReply(data)
}
