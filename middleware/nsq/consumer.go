package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/wansnow/calculation_server/config"
	"log"
	"sync"
)

func StartNewConsumer(topic, channel string, handler nsq.Handler, waitGroup *sync.WaitGroup, stop chan int) {
	waitGroup.Add(1)
	waitGroup.Done()
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
	<-stop
	consumer.Stop()
}
