package process

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"req_resp/conf"
)

func GetCodeArr()(arr []string){
	//url := "http://61.129.249.241:8880/stockdict/full?fmt=j&ver=0&mid=150"
	url := conf.GetConfig().StockDictUrl

	r, err := http.DefaultClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	body, _ := ioutil.ReadAll(r.Body)
	buffer :=bytes.NewBuffer(body)
	js,_:=simplejson.NewFromReader(buffer)
	tableArr, _ := js.Get("list").Array()
	fmt.Printf("一共有%d行全量码表信息\n", len(tableArr))

	var codeArr []string
	for i := range tableArr {
		table := js.Get("list").GetIndex(i)
		code := table.Get("c").MustString()
		codeArr = append(codeArr, code)
		//fmt.Println(code)
	}
	//fmt.Println(codeArr)
	return codeArr
}