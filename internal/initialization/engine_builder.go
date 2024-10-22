package initialization

import (
	"errors"
	"github.com/passsquale/key-value-storage/internal/configuration"
	"github.com/passsquale/key-value-storage/internal/database/storage"
	"github.com/passsquale/key-value-storage/internal/database/storage/engine/in_memory"
	"go.uber.org/zap"
)

const (
	InMemoryEngine = "in_memory"
)

var supportedEngineTypes = map[string]struct{}{
	InMemoryEngine: {},
}

const defaultPartitionsNumber = 10

func CreateEngine(cfg *configuration.EngineConfig, logger *zap.Logger) (storage.Engine, error) {
	if cfg == nil {
		return in_memory.NewEngine(in_memory.HashTableBuilder, defaultPartitionsNumber, logger)
	}

	if cfg.Type != "" {
		_, found := supportedEngineTypes[cfg.Type]
		if !found {
			return nil, errors.New("engine type is incorrect")
		}
	}

	return in_memory.NewEngine(in_memory.HashTableBuilder, defaultPartitionsNumber, logger)
}
