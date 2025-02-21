package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/passsquale/key-value-storage/internal/tools"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
	"time"
)

type TCPHandler = func(context.Context, []byte) []byte

type TCPServer struct {
	address     string
	semaphore   tools.Semaphore
	idleTimeout time.Duration
	messageSize int
	logger      *zap.Logger
}

func NewTCPServer(
	address string,
	maxConnectionsNumber int,
	maxMessageSize int,
	idleTimeout time.Duration,
	logger *zap.Logger,
) (*TCPServer, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	if maxConnectionsNumber <= 0 {
		return nil, errors.New("invalid number of max connections")
	}

	return &TCPServer{
		address:     address,
		semaphore:   tools.NewSemaphore(maxConnectionsNumber),
		idleTimeout: idleTimeout,
		messageSize: maxMessageSize,
		logger:      logger,
	}, nil
}

func (s *TCPServer) HandleQueries(ctx context.Context, handler TCPHandler) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			connection, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}

				s.logger.Error("failed to accept", zap.Error(err))
				continue
			}

			wg.Add(1)
			go func(connection net.Conn) {
				s.semaphore.Acquire()

				defer func() {
					s.semaphore.Release()
					wg.Done()
				}()

				s.handleConnection(ctx, connection, handler)
			}(connection)
		}
	}()

	go func() {
		defer wg.Done()

		<-ctx.Done()
		if err := listener.Close(); err != nil {
			s.logger.Warn("failed to close listener", zap.Error(err))
		}
	}()

	wg.Wait()
	return nil
}

func (s *TCPServer) handleConnection(ctx context.Context, connection net.Conn, handler TCPHandler) {
	request := make([]byte, s.messageSize)

	for {
		if err := connection.SetDeadline(time.Now().Add(s.idleTimeout)); err != nil {
			s.logger.Warn("failed to set read deadline", zap.Error(err))
			break
		}

		count, err := connection.Read(request)
		if err != nil {
			if err != io.EOF {
				s.logger.Warn("failed to read", zap.Error(err))
			}

			break
		}

		response := handler(ctx, request[:count])
		if _, err := connection.Write(response); err != nil {
			s.logger.Warn("failed to write", zap.Error(err))
			break
		}
	}

	if err := connection.Close(); err != nil {
		s.logger.Warn("failed to close connection", zap.Error(err))
	}
}
