package start_game

import (
	"fmt"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/server/calculation_server/service/calculation_service"
	"github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestStartGame(t *testing.T) {
	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	projectPath := strings.Split(filePath, "test")[0]
	err = config.LoadConfig(projectPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	id := "test"
	pyFunc := `
move("%s", 2)`
	go func() {
		err := logic_service.StartLogic(id, fmt.Sprintf(pyFunc, id))
		if err != nil {

		}
	}()

	time.Sleep(1)

	go calculation_service.StartGameCalculation(id)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
