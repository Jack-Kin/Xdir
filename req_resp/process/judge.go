package process

import (
	"fmt"
	"req_resp/proto/msg6030"
	"req_resp/proto/msg6131"
)

func Judge6131(resMsg *msg6131.ResponseMessage, code *string) {

	// 得到日期数组和value数组
	var dateArr []uint32
	var value1Arr []int32
	var value2Arr []int32
	for _, x := range resMsg.Data {
		//fmt.Printf("\t%d\t\t%d\t\t%d\t\t%d\n", i, x.Tradeday, x.Value1, x.Value2)
		dateArr = append(dateArr, x.Tradeday)
		value1Arr = append(value1Arr,x.Value1)
		value2Arr = append(value2Arr,x.Value2)
	}
	dateResult := DateAscend(&dateArr)
	if dateResult == false {
		fmt.Printf("code%s, 日期出问题\n", *code)
	}
	if (len(value1Arr) == 1 && value1Arr[0] == 0) || (len(value2Arr) == 1 && value2Arr[0] == 0){
		fmt.Printf("code%s, value为空\n", *code)
	}
	value1Result := ValueJudge(&value1Arr)
	value2Result := ValueJudge(&value2Arr)
	if value1Result == false || value2Result == false {
		fmt.Printf("code%s, value值错误\n", *code)
	}
}


/*
对message6030的检测逻辑：
  - 检验TradeDate是否升序
  - 检验开高低收是否有=0的,有就直接返回
  - 检验开高低收是否有大幅波动
  - 判断是否高是最大的
  - 判断是否低是最小的
 */

func Judge6030(resMsg *msg6030.ResponseMessage, code *string) {

	var dateArr []uint32
	var openArr []int32
	var highArr []int32
	var lowArr []int32
	var closeArr []int32
	for _, x := range resMsg.Data5 {
		dateArr = append(dateArr, x.Tradeday)
		openArr = append(openArr, x.Open)
		highArr = append(highArr, x.High)
		lowArr = append(lowArr, x.Low)
		closeArr = append(closeArr, x.Close)
	}
	//// 判断数据是否空......退市的太多了，不print了
	//if len(dateArr) + len(openArr) + len(highArr) + len(lowArr) + len(closeArr) == 0 {
	//	fmt.Printf("code %s, 数据为空\n", *code)
	//	return
	//}

	// 判断日期是否递增
	dateResult := DateAscend(&dateArr)
	if dateResult == false {
		fmt.Printf("code %s, 日期出问题\n", *code)
		return
	}
	// 判断数据是否对齐
	if len(openArr) != len(dateArr) || len(highArr) != len(dateArr) || len(lowArr) != len(dateArr) || len(closeArr) != len(dateArr) {
		fmt.Printf("code %s, 数据不对齐\n", *code)
		return
	}
	for i := 0; i < len(dateArr); i++ {
		// 判断开高低收是否有0
		if lowArr[i] == 0 || openArr[i] == 0 || highArr[i] == 0 || closeArr[i] == 0 {
			fmt.Printf("code %s, date %d, 开高低收有值为0\n", *code, dateArr[i])
			break
		}
		// 判断开高低收是否有值大于设定的最大界限
		var MaxValue int32 = 1000000000
		if highArr[i] > MaxValue {
			fmt.Printf("code %s, date %d, 开高低收有值太大\n", *code, dateArr[i])
			return
		}
		// 判断高是最大的
		if openArr[i] > highArr[i] || lowArr[i] > highArr[i] || lowArr[i] > highArr[i] {
			fmt.Printf("code %s, date %d, high不是最大值\n", *code, dateArr[i])
			break
		}
		// 判断低是最小的
		if openArr[i] < lowArr[i] || highArr[i] < lowArr[i] || lowArr[i] < lowArr[i] {
			fmt.Printf("code %s, date %d, low不是最小值\n", *code, dateArr[i])
			break
		}
	}
}