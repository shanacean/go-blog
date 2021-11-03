package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {

	filePath := "log/log.log"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	writers := []io.Writer{
		file,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.SetOutput(fileAndStdoutWriter)

	// log分割
	logs, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logs,
		logrus.DebugLevel: logs,
		logrus.FatalLevel: logs,
		logrus.ErrorLevel: logs,
		logrus.WarnLevel:  logs,
		logrus.PanicLevel: logs,
	}

	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()

		stopTime := time.Since(startTime)

		duration := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))

		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"Duration":  duration,
			"Method":    method,
			"DataSize":  dataSize,
			"Ip":        clientIp,
			"UserAgent": userAgent,
			"Path":      path,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
