package middleware

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"login/app/config"
	"os"
	"path"
	"strings"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	conf := config.Conf
	logFilePath := conf.Log.DirName
	logFileName := conf.Log.FileName + ".log"

	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("open log file failed, error: %s\n", err)

	}
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	if conf.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		// 设置输出
		logger.Out = src
		logger.SetLevel(logrus.InfoLevel)
		// 设置 rotatelogs
		logWriter, err := rotatelogs.New(
			// 分割后的文件名称
			strings.Replace(fileName, ".log", "", -1)+".%Y-%m-%d.log",
			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(fileName),
			// 设置最大保存时间(7天)
			rotatelogs.WithMaxAge(conf.Log.MaxAge*24*time.Hour),
			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			log.Fatalf("rotatelogs failed, error: %s\n", err)
		}
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		lfsHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		// 新增 Hook
		logger.AddHook(lfsHook)
	}

	return func(c *gin.Context) {
		reqUri := c.Request.RequestURI
		if reqUri == "/favicon.ico" {
			return
		}
		startTime := time.Now()
		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		if conf.Log.Debug {
			now := time.Now().Format("2006-01-02 15:04:05")
			logger.Infof("%s | %3d | %13v | %15s | %s  %s",
				now,
				statusCode,
				latencyTime,
				clientIP,
				reqMethod,
				reqUri,
			)
		} else {
			logger.WithFields(logrus.Fields{
				"status_code":  statusCode,
				"latency_time": latencyTime,
				"client_ip":    clientIP,
				"req_method":   reqMethod,
				"req_uri":      reqUri,
			}).Info()
		}
	}
}

// 日志记录到 Mysql
func LoggerToMysql() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

