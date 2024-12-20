package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutputFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args...)
}
func Info(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "info")
	logrus.WithFields(fields).Info(args...)
}
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "warn")
	logrus.WithFields(fields).Warn(args...)
}
func Error(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "error")
	logrus.WithFields(fields).Error(args...)
}
func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "fatal")
	logrus.WithFields(fields).Fatal(args...)
}

func setOutputFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.Mkdir("./runtime/log", os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("create log dir '%s' failed, error: %s", "./runtime/log", err))
		}
	}
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", logName+"_"+timeStr+".log")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open log file failed, err ", err)
		return
	}
	logrus.SetOutput(file)
	logrus.SetLevel(level)
}

func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.Mkdir("./runtime/log", os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("create log dir '%s' failed, error: %s", "./runtime/log", err))
		}
	}
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", "success_"+timeStr+".log")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("create log dir '%s' failed, error: %s", "./runtime/log", err))
	}
	conf := gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d\" %s %s\n",
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.ErrorMessage)
		},
		Output: io.MultiWriter(os.Stdout, file),
	}
	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				errDir = os.MkdirAll("./runtime/log", os.ModePerm)
				if errDir != nil {
					panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", errDir))
				}

				timeStr := time.Now().Format("2006-01-02")
				fileName := path.Join("./runtime/log", "error_"+timeStr+".log")

				f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
				if errFile != nil {
					fmt.Println("open log file failed, err ", errFile)
					return
				}

				timeNowStr := time.Now().Format("2006-01-02 15:04:05")
				f.WriteString("panic at " + timeNowStr + "\n")
				f.WriteString(fmt.Sprintf("err: %v", err))
				f.WriteString("stacktrace from " + string(debug.Stack()) + "\n")
				f.WriteString("---------------------------------\n")
				f.Close()

				c.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": fmt.Sprintf("err: %v", err),
				})
				c.Abort()
			}
		}
	}()
	c.Next()
}
