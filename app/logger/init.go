package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func Init() error {
	debug := viper.GetBool("app.debug")
	rotateWriter := &lumberjack.Logger{
		Filename:   viper.GetString("log.path"),
		MaxSize:    viper.GetInt("log.size"),
		MaxAge:     viper.GetInt("log.age"),
		MaxBackups: viper.GetInt("log.backup"),
		LocalTime:  true,
		Compress:   viper.GetBool("log.compress"),
	}

	writes := []io.Writer{rotateWriter}
	if debug {
		writes = append(writes, os.Stdout)
	}

	logrus.SetOutput(io.MultiWriter(writes...))
	logrus.SetReportCaller(debug)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   "2006-01-02 15:04:05",
	})

	return nil
}
