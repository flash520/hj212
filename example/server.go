/**
 * @Author: koulei
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2023/3/7 22:13
 */

package main

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/flash520/hj212"
	"github.com/flash520/hj212/consts"
)

func init() {
	// 初始化日志库
	if err := SetLogLevel(InfoLevel); err != nil {
		panic(err)
	}
	log.StandardLogger().Formatter = &prefixed.TextFormatter{
		ForceColors:      true,
		ForceFormatting:  true,
		FullTimestamp:    true,
		DisableTimestamp: true,
		QuoteEmptyFields: true,
		DisableSorting:   true,
		TimestampFormat:  "2006-01-02 15:04:05",
	}

	log.SetReportCaller(true)
	// log.SetFormatter(&log.TextFormatter{
	// 	FullTimestamp: true,
	// 	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
	// 		filename := path.Base(f.File)
	// 		return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
	// 	},
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
}

// Level 日志等级
type Level string

var (
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
	InfoLevel  Level = "info"
	DebugLevel Level = "debug"
	TraceLevel Level = "trace"
)

// SetLogLevel 设置日志级别
func SetLogLevel(level Level) error {
	lv, err := log.ParseLevel(string(level))
	if err != nil {
		return err
	}
	log.StandardLogger().SetLevel(lv)
	return nil
}

func main() {
	server := hj212.NewServer(hj212.Option{
		ListenAddress: "0.0.0.0:8192",
		SendChanSize:  0,
		Keepalive:     0,
		CloseHandler:  nil,
	})

	if err := server.Run(); err != nil {
		log.Error(consts.ServerName, err.Error())
	}
}
