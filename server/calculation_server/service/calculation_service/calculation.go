package calculation_service

import (
	"github.com/nsqio/go-nsq"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
)

type CalculationMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *CalculationMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	_ = func_msg.Decode(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}
