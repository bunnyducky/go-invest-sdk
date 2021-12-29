package investsdk

import (
	"fmt"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConf struct {
	Level   zapcore.Level //-1: debug, 0:info, 1:warn ...
	File    string        //default log file, if empty, log to stdout
	ErrFile string        //write err only if not empty
}

//write error to individual file
//colored log
func NewLogger(conf ZapConf, defaultW io.Writer) (*zap.Logger, func(), error) {
	var cores []zapcore.Core

	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder //2020-10-29T11:59:18.606+0800
	{
		//default logger
		var defaultWriter zapcore.WriteSyncer
		if conf.File != "" {
			f, err := os.OpenFile(conf.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
				return nil, nil, err
			}
			defaultWriter = zapcore.AddSync(f)
		} else {
			defaultWriter = zapcore.AddSync(defaultW)
		}
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf), defaultWriter, conf.Level))
	}

	if conf.ErrFile != "" { //error logger
		f, err := os.OpenFile(conf.ErrFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return nil, nil, err
		}
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf), zapcore.AddSync(f), zap.WarnLevel))
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	recoverZapGlobal := zap.ReplaceGlobals(logger)
	recoverStd := zap.RedirectStdLog(logger)

	return logger, func() {
		recoverZapGlobal()
		recoverStd()
		logger.Sync()
	}, nil
}

type CronLogger zap.SugaredLogger

func (logger *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	((*zap.SugaredLogger)(logger)).Infow(msg, keysAndValues...)
}
func (logger *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	((*zap.SugaredLogger)(logger)).Errorw(fmt.Sprintf("%s, err: %v", msg, err), keysAndValues...)
}
