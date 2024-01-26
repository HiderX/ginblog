package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	latestLogPath := "log/latest_log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(latestLogPath),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	//logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		spendTime := int(endTime.Sub(startTime).Milliseconds())
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := max(c.Writer.Size(), 0)
		method := c.Request.Method
		path := c.Request.URL.Path
		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostName,
			"status":     statusCode,
			"latency":    strconv.Itoa(spendTime) + "ms",
			"ip":         clientIP,
			"method":     method,
			"path":       path,
			"data_size":  dataSize,
			"user_agent": userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
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
