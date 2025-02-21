package initialization

import (
	"errors"
	"github.com/passsquale/key-value-storage/internal/configuration"
	"github.com/passsquale/key-value-storage/internal/database/storage"
	"github.com/passsquale/key-value-storage/internal/database/storage/wal"
	"github.com/passsquale/key-value-storage/internal/tools"
	"go.uber.org/zap"
	"time"
)

const defaultFlushingBatchSize = 100
const defaultFlushingBatchTimeout = time.Millisecond * 10
const defaultMaxSegmentSize = 10 << 20
const defaultWALDataDirectory = "./data/kv-storage/wal"

func CreateWAL(cfg *configuration.WALConfig, logger *zap.Logger) (storage.WAL, error) {
	flushingBatchSize := defaultFlushingBatchSize
	flushingBatchTimeout := defaultFlushingBatchTimeout
	maxSegmentSize := defaultMaxSegmentSize
	dataDirectory := defaultWALDataDirectory

	if cfg != nil {
		if cfg.FlushingBatchLength != 0 {
			flushingBatchSize = cfg.FlushingBatchLength
		}

		if cfg.FlushingBatchTimeout != 0 {
			flushingBatchTimeout = cfg.FlushingBatchTimeout
		}

		if cfg.MaxSegmentSize != "" {
			size, err := tools.ParseSize(cfg.MaxSegmentSize)
			if err != nil {
				return nil, errors.New("max segment size is incorrect")
			}

			maxSegmentSize = size
		}

		if cfg.DataDirectory != "" {
			dataDirectory = cfg.DataDirectory
		}

		fsReader := wal.NewFSReader(dataDirectory, logger)
		fsWriter := wal.NewFSWriter(dataDirectory, maxSegmentSize, logger)
		return wal.NewWAL(fsWriter, fsReader, flushingBatchTimeout, flushingBatchSize), nil
	} else {
		return nil, nil
	}
}
