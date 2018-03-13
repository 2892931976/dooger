package main

import (
	"log"
	"net/http"
	"bytes"

)

type Ding struct {
	WebHookUrl string
}

func (d *Ding) Write(p []byte) (n int, err error) {
	//remove last \n byte
	p = p[:len(p)-2]
	buf := bytes.NewBufferString(`{"msgtype":"text","text":{"content":"`)
	buf.Write(p)
	buf.WriteString(`"}}`)
	go http.Post(d.WebHookUrl, "application/json", buf)
	return 0, err
}


var lger *log.Logger

func init() {
	dingDingTalker :=  Ding{"https://oapi.dingtalk.com/robot/send?access_token=e28e8b2efdd05a9954f888ab16b2e059706628f85590e269cae996eae7fbbf8f"}
	lger = log.New(&dingDingTalker, "eric:", log.Lshortfile | log.LstdFlags)
}

func main() {

	lger.Println("test test:awesome awesome")
	//lger.Println("test test:twoww")

	select {}

}
