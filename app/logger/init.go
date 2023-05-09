package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() error {
	rotateWriter := &lumberjack.Logger{
		Filename:   viper.GetString("log.path"),
		MaxSize:    viper.GetInt("log.size"),
		MaxAge:     viper.GetInt("log.age"),
		MaxBackups: viper.GetInt("log.backup"),
		LocalTime:  true,
		Compress:   viper.GetBool("log.compress"),
	}

	debug := viper.GetBool("app.debug")
	writers := []io.Writer{rotateWriter}

	if debug {
		writers = append(writers, os.Stdout)
	}

	logrus.SetOutput(io.MultiWriter(writers...))
	logrus.SetReportCaller(debug)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   "2006-01-02 15:04:05",
	})

	return nil
}
