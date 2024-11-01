package middleware

import (
	"fmt"
	"go-blog-server/utils"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Log() gin.HandlerFunc {
	// 当前程序执行的目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前程序执行目录失败")
	}
	filePath := fmt.Sprintf("%s/%s/%s", dir, utils.LogPath, utils.LogFile) // 日志目录文件
	//linkName := utils.LinkName // 软连接到日志文件

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开日志文件错误", err)
	}

	logger := logrus.New()             // 构造函数，构造一个*Logger
	logger.Out = f                     // 将日志输出到文件
	logger.SetLevel(logrus.DebugLevel) // 设置日志记录的级别

	// 日志分割
	logWriter, _ := retalog.New( // 构造函数
		filePath+"%Y%m%d.log",                  // 参数1 日志文件名称
		retalog.WithMaxAge(7*24*time.Hour),     // 参数2 文件最大保存时间
		retalog.WithRotationTime(24*time.Hour), // 参数3 日志切割时间间隔
		//retalog.WithLinkName(linkName),         // 参数4 生成链接，指向最新日志文件
	)
	// 所有级别的日志写入到logWriter
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

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) // 函数耗时
		hostName, err := os.Hostname()                                                               // 主机名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    // 请求状态码
		clientIp := c.ClientIP()           // 客户端IP
		userAgent := c.Request.UserAgent() // 浏览器
		dataSize := c.Writer.Size()        // 请求大小
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method   // 请求方法
		path := c.Request.RequestURI // 请求路径

		// logrus的标准写法
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		// gin框架内部错误
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		// 状态码错误
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
