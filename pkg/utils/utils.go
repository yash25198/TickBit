package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func Retry(logger *zap.Logger, ctx context.Context, dur time.Duration, f func() error) error {
	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	err := f()
	for err != nil {
		// Skip retrying if it's a `NoRetryError`
		var e *NoRetryError
		if errors.As(err, &e) {
			return err
		}

		logger.Debug("retrying", zap.Any("error", err.Error()))
		select {
		case <-ctx.Done():
			return fmt.Errorf("%v: %v", ctx.Err(), err)
		case <-ticker.C:
			err = f()
		}
	}
	return nil
}

type NoRetryError struct {
	err error
}

func NewNoRetryError(err error) *NoRetryError {
	return &NoRetryError{
		err: err,
	}
}

func (err *NoRetryError) Error() string {
	if err.err != nil {
		return err.err.Error()
	}
	return ""
}

func (err *NoRetryError) Unwrap() error {
	return err.err
}
