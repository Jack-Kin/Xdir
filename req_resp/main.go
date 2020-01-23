package main

import (
	"fmt"
	"req_resp/client"
	"req_resp/process"
	_ "req_resp/proto"
	"runtime"
	"time"
)


func main() {
	/*
	看说明文档
	*/
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量


	//// 正式
	//codeArr := process.GetCodeArr()
	//for _,code := range codeArr {
	//	client.GetData(code)
	//}


	// 测试时间性能
	t1 := time.Now()
	codeArr := process.GetCodeArr()
	t2 := time.Now()
	for i,code := range codeArr {
		if code == "430025" {
			client.GetData(code)
		}
		if (i+1) % 1000 == 0 {
			fmt.Printf("跑了%d个\n",i+1)
		}
		//client.GetData(code)
	}
	t3 := time.Now()
	fmt.Println("读取code费时：",t2.Sub(t1))
	fmt.Println("检测code费时：", t3.Sub(t2))


	//// 测试返回包
	////client.GetData(codeArr[0])
	//client.GetData("SR005")
	//client.GetData("300059")
	//client.GetData("688299")


	fmt.Printf("结束\n")
}

