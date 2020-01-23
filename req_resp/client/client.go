package client

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"req_resp/conf"
	"req_resp/process"
	"req_resp/proto"
	"req_resp/proto/header"
	"req_resp/proto/msg6030"
	//"req_resp/proto/msg6131"
)

//发送信息
func sender(conn net.Conn, code string) {
	//resquest message


	//// msg6030
	reqMsg := msg6030.NewRequest(5, 0,
		0,0, 7, 2,
		0,0, 0,100,0, []byte(code))
	reqHead := header.New(6030, reqMsg.Size())
	resHeader := new(header.Header)
	resMsg := msg6030.NewResponse()

	////// msg6131
	//reqMsg := msg6131.NewRequest(1, 0, 0, 0, 999, 7, 1792, 150, []byte(code))
	//reqHead := header.New(6131, reqMsg.Size())
	//resHeader := new(header.Header)
	//resMsg := msg6131.NewResponse()


	//tail
	tail := proto.NewTail()


	//fmt.Println("\tlocal address is:", conn.LocalAddr())
	//fmt.Println("\t\tH:", reqHead)
	//fmt.Println("\t\tM:", reqMsg)

	buf := new(bytes.Buffer)
	encoder := proto.NewEncoder(buf)
	encoder.Push(reqHead).Push(reqMsg).Push(tail)
	err := encoder.Error()
	if err != nil {
		fmt.Println("encode err:", err)
	}
	//fmt.Printf("\t% d", buf.Bytes())

	// 向conn写入数据
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	decoder := proto.NewDecoder(conn)

	decoder.Fetch(resHeader).Fetch(resMsg)
	//fmt.Println(resMsg.Data5)


	// 数据有错误在这个函数里处理并输出, 判断函数名字和msgid有关
	//process.Judge6131(resMsg, &code)
	process.Judge6030(resMsg, &code)



	err = decoder.Error()
	if err != nil {
		fmt.Println("decode  err:", err)
		Log(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return
	}


	//fmt.Printf("*** begin code %s:\n", code)
	//fmt.Println("\tlocal address is:", conn.LocalAddr())
	//fmt.Println("\t\tH:", reqHead)
	//fmt.Println("\t\tM:", reqMsg)
	//fmt.Println("\tremote address is:", conn.RemoteAddr())
	//fmt.Println("\t\tH:", resHeader)
	//fmt.Println("\t\tM:", resMsg)
}

//日志
func Log(v ...interface{}) {
	log.Println(v...)
}

func GetData(code string) {
	//server := "43.254.147.74:80" 现在改为从配置文件中读取server
	server := conf.GetConfig().AppServerAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	// fmt.Println("connection success")
	sender(conn, code)
}
