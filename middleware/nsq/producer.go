package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/wansnow/calculation_server/config"
	"log"
	"sync"
)

func StartNewProducer(gameId string, msgChan chan []byte, waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	defer waitGroup.Done()
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(config.NsqC.NsqdUrl, cfg)
	if err != nil {
		log.Println(err)
	}
	// 发布消息
	for {
		msg := <-msgChan
		if string(msg) == "end" {
			break
		}
		if err := producer.Publish(fmt.Sprintf("topic_%s", gameId), msg); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
	}
	producer.Stop()

}
