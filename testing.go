package investsdk

import (
	"net/http"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ExecuteInfrastructure struct {
	Client *Client
	Logger *zap.SugaredLogger
}

type TestOut struct{ *testing.T }

func (to *TestOut) Write(p []byte) (n int, err error) {
	to.T.Log(string(p))
	return len(p), nil
}

func GetTestInfra(t *testing.T) *ExecuteInfrastructure {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&TestOut{t}),
		zapcore.DebugLevel,
	)
	logger := zap.New(zapcore.NewTee(core), zap.AddCaller())
	recoverZapGlobal := zap.ReplaceGlobals(logger)
	recoverStd := zap.RedirectStdLog(logger)
	t.Cleanup(func() {
		recoverZapGlobal()
		recoverStd()
	})
	sLogger := logger.Sugar()

	investSDKClient := Client{
		Host: "http://127.0.0.1:3500",
		// Host:   "https://invest-api.partyparrot.finance",
		Client: http.DefaultClient,
		Logger: sLogger,
	}

	return &ExecuteInfrastructure{
		Client: &investSDKClient,
		Logger: sLogger,
	}

}
