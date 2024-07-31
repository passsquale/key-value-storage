package initialization

import (
	"errors"
	"github.com/evgeniy-roslyackov/key-value-storage/internal/configuration"
	"github.com/evgeniy-roslyackov/key-value-storage/internal/network"
	"github.com/evgeniy-roslyackov/key-value-storage/internal/tools"
	"go.uber.org/zap"
	"time"
)

const defaultServerAddress = "localhost:8080"
const defaultMaxConnectionNumber = 100
const defaultMaxMessageSize = 2048
const defaultIdleTimeout = time.Minute * 5

func CreateNetwork(cfg *configuration.NetworkConfig, logger *zap.Logger) (*network.TCPServer, error) {
	address := defaultServerAddress
	maxConnectionsNumber := defaultMaxConnectionNumber
	maxMessageSize := defaultMaxMessageSize
	idleTimeout := defaultIdleTimeout

	if cfg != nil {
		if cfg.Address != "" {
			address = cfg.Address
		}

		if cfg.MaxConnections != 0 {
			maxConnectionsNumber = cfg.MaxConnections
		}

		if cfg.MaxMessageSize != "" {
			size, err := tools.ParseSize(cfg.MaxMessageSize)
			if err != nil {
				return nil, errors.New("incorrect max message size")
			}

			maxMessageSize = size
		}

		if cfg.IdleTimeout != 0 {
			idleTimeout = cfg.IdleTimeout
		}
	}

	return network.NewTCPServer(address, maxConnectionsNumber, maxMessageSize, idleTimeout, logger)
}
