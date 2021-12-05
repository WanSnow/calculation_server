package calculation_service

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	middleware_nsq "github.com/wansnow/calculation_server/middleware/nsq"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
	"log"
)

func StartGameCalculation(id string) {
	middleware_nsq.StartNewConsumer(fmt.Sprintf("topic_%s", id), "main", nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(func_msg.Decode(message.Body))
		return nil
	}))
}
