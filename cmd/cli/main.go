package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/evgeniy-roslyackov/key-value-storage/internal/network"
	"github.com/evgeniy-roslyackov/key-value-storage/internal/tools"
	"go.uber.org/zap"
	"os"
	"syscall"
	"time"
)

func main() {
	address := flag.String("address", "localhost:8080", "Address of the kv-storage")
	idleTimeout := flag.Duration("idle_timeout", time.Minute, "Idle timeout for connection")
	maxMessageSizeStr := flag.String("max_message_size", "4KB", "Max message size for connection")
	flag.Parse()

	logger, _ := zap.NewProduction()
	maxMessageSize, err := tools.ParseSize(*maxMessageSizeStr)
	if err != nil {
		logger.Fatal("failed to parse max message size", zap.Error(err))
	}

	reader := bufio.NewReader(os.Stdin)
	client, err := network.NewTCPClient(*address, maxMessageSize, *idleTimeout)
	if err != nil {
		logger.Fatal("failed to connect with server", zap.Error(err))
	}

	for {
		fmt.Print("[kv-storage] > ")
		request, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				logger.Fatal("connection was closed", zap.Error(err))
			}

			logger.Error("failed to read user query", zap.Error(err))
		}

		response, err := client.Send([]byte(request))
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				logger.Fatal("connection was closed", zap.Error(err))
			}

			logger.Error("failed to send query", zap.Error(err))
		}

		fmt.Println(string(response))
	}
}
