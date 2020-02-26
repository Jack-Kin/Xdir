package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	//发送消息使用导的url
	sendurl = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	//获取token使用导的url
	get_token = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

var requestError = errors.New("request error,check url or network")

type access_token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

//定义一个简单的文本消息格式
type send_msg struct {
	Touser  string            `json:"touser"`
	Toparty string            `json:"toparty"`
	Totag   string            `json:"totag"`
	Msgtype string            `json:"msgtype"`
	Agentid int               `json:"agentid"`
	Text    map[string]string `json:"text"`
	Safe    int               `json:"safe"`
}

type send_msg_error struct {
	Errcode int    `json:"errcode`
	Errmsg  string `json:"errmsg"`
}


func main(){
	touser := flag.String("t", "@all", "-t user 直接接收消息的用户昵称")
	agentid := flag.Int("i", xxx, "-i 0 指定agentid")
	corpid := flag.String("p", "xxx", "-p corpid 必须指定")
	corpsecret := flag.String("s", "xxx", "-s corpsecret 必须指定")
	flag.Parse()

	if *corpid == "" || *corpsecret == "" {
		flag.Usage()
		return
	}

	// 获取token
	token, err := Get_token(*corpid, *corpsecret)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("获取token成功：", token)

	// windows/linux环境下
	path, _ := os.Getwd()
	// ide下
	//path += "/req_resp/"
	//println(path)


	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// 对*.txt一个个发送
	for _, file := range files {
		length := len(file.Name())
		if file.Name()[length - 4:] == ".txt" {
			println(file.Name())
			// 读取文件
			fi, err := os.Open(path + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			defer fi.Close()
			msg, err := ioutil.ReadAll(fi)
			if err != nil {
				fmt.Println("read to fd fail", err)
				msg = []byte("")
			}
			//if len(msg) == 0{
			//	str := "空"
			//	msg = ([]byte)(str)
			//}

			var m send_msg = send_msg{
				Touser: *touser,
				Toparty: "",
				Totag: "",
				Msgtype: "text",
				Agentid: *agentid,
				Text: map[string]string{"content": string(msg)},
				Safe: 0,
			}
			// json编组
			buf, err := json.Marshal(m)
			if err != nil {
				return
			}
			// 发送消息
			err = Send_msg(token.Access_token, buf)
			if err != nil {
				println(err.Error())
			} else {
				fmt.Println("发送消息成功", string(buf))
			}
		}
	}
}

//发送消息.msgbody 必须是 API支持的类型
func Send_msg(Access_token string, msgbody []byte) error {
	body := bytes.NewBuffer(msgbody)
	resp, err := http.Post(sendurl+Access_token, "application/json", body)
	if resp.StatusCode != 200 {
		return requestError
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var e send_msg_error
	err = json.Unmarshal(buf, &e)
	if err != nil {
		return err
	}
	if e.Errcode != 0 && e.Errmsg != "ok" {
		return errors.New(string(buf))
	}
	return nil
}

//通过corpid 和 corpsecret 获取token
func Get_token(corpid, corpsecret string) (at access_token, err error) {
	resp, err := http.Get(get_token + corpid + "&corpsecret=" + corpsecret)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = requestError
		return
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &at)
	if at.Access_token == "" {
		err = errors.New("corpid or corpsecret error.")
	}
	return
}
