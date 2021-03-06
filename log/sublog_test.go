package log

import (
	"context"
	"errors"
	"testing"

	"github.com/kitabisa/perkakas/v2/ctxkeys"
)

// TODO: TO BE DEPRECATED after v2.14.6
func TestSublogger(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxkeys.CtxXKtbsRequestID, "111111")

	subLog := GetSublogger(ctx, "test-context-name-1")
	subLog.Err(errors.New("test-error")).Msg("test-message")

	subLog2 := GetSublogger(ctx, "test-context-name-2")
	subLog2.Info().Msg("test-message")

	// as log is not printed when test is success, need to uncomment line below :D
	// t.FailNow()
}
