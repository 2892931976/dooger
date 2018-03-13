package dogger

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

var logger *log.Logger

func NewLogger(hookUrl string) *log.Logger {
	dingDingTalker := Ding{hookUrl}
	logger = log.New(&dingDingTalker, "eric:", log.Lshortfile | log.LstdFlags)
	return logger
}

func Println(v ...interface{}) {
	if logger == nil {
		log.Fatal("请设置钉钉bot Web Hook Url")
	}
	logger.Println(v)
}
