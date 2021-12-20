package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/wansnow/calculation_server/config"
	"log"
)

func StartNewProducer(topic string, msgChan chan []byte) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(config.NsqC.NsqdUrl, cfg)
	if err != nil {
		log.Println(err)
	}
	// 发布消息
	for {
		msg := <-msgChan
		if err := producer.Publish(topic, msg); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
	}
}
