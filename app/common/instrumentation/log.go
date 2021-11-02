package instrumentation

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/common/shared"
)

// LogFromContext add request_id and operation in to log entries
func LogFromContext(ctx context.Context, log zerolog.Logger, operation string) (*zerolog.Logger, error) {
	xRequestID, ok := ctx.Value(shared.KeyRequestID).(string)
	if !ok {
		return nil, fmt.Errorf("invalid request_id")
	}

	log = log.With().Str("operation", operation).Logger()
	log = log.With().Str("request_id", xRequestID).Logger()

	return &log, nil
}
