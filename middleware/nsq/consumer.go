package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/wansnow/calculation_server/config"
	"log"
)

func StartNewConsumer(topic, channel string, handler nsq.Handler) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(handler)
	// nsqlookupd
	//[]string
	if err := consumer.ConnectToNSQD(config.NsqC.NsqdUrl); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
	consumer.Stop()
}
