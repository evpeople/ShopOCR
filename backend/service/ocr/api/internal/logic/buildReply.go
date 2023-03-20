package logic

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/evpeople/ShopOCR/common/errorx"
	"github.com/evpeople/ShopOCR/service/ocr/api/internal/types"
)

type PairData [4][2]int

// 获取Paris配对
func findMatchingBrackets(data string, symbolBegin, symbolEnd rune) ([][]int, error) {

	var stack []int
	var pairs [][]int

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
	pairs, err := findMatchingBrackets(data, '(', ')')
	if err != nil {
		errorx.NewDefaultError(err.Error())
	}
	pairs2, err := findMatchingBrackets(data, '[', ']')
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
