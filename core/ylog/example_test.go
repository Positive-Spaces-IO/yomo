package ylog_test

import (
	"io"
	"net"

	"github.com/yomorun/yomo/core/ylog"
)

func Example() {
	// text format logger
	logger := ylog.NewFromConfig(ylog.Config{
		Level:       "warn",
		Format:      "text",
		ErrorOutput: "stdout",
		DisableTime: true,
	})

	ylog.SetDefault(logger.With("hello", "yomo").WithGroup("syslog"))

	ylog.Debug("debug", "aaa", "bbb")
	ylog.Info("info", "ccc", "ddd")
	ylog.Warn("warn", "eee", "fff")
	ylog.Error("error", io.EOF, "eee", "fff")

	// json format logger
	sysLogger := ylog.NewFromConfig(ylog.Config{
		Level:       "error",
		Format:      "json",
		ErrorOutput: "stdout",
		DisableTime: true,
	})

	sysLogger = sysLogger.WithGroup("syslog")

	sysLogger.Error("sys error", net.ErrClosed, "ggg", "hhh")

	// Output:
	// level=WARN msg=warn hello=yomo syslog.eee=fff
	// level=ERROR msg=error hello=yomo syslog.eee=fff syslog.err=EOF
	// {"level":"ERROR","msg":"sys error","syslog":{"ggg":"hhh","err":"use of closed network connection"}}
}
