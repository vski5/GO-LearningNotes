package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

//WriterSyncer ：指定日志将写到哪里去。
func getLogWriter() zapcore.WriteSyncer {
	//创造
	os.Create("./test.log")
	//打开创造的文件，追加日志
	file, _ := os.OpenFile("./test.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend, 0744)
	return zapcore.AddSync(file) //为了满足接口，转格式。
}

var logger *zap.Logger

//InitLogger 初始化logger
func InitLogger() {
	//上文用到的默认的配置
	//logger, _ = zap.NewProduction()

	//设定三个参数：编码器、日志写入目的地、日志级别。
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//自定义配置
	logger = zap.New(core)

	//logger.Info("hello world") //输出文字"hello world"到文件中
}

//其他的手段：
//更改时间编码,属于Encoder:编码器(如何写入日志)的一部分
func getTimeEncoder() zapcore.Encoder {
	//zap.NewProductionEncoderConfig()本质上就是造了一个zapcore.EncoderConfig结构体，我也可以造一个自己的结构体，然后设置好我想要的配置。
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//zapcore.NewTimeEncoder作用是把时间转换成字符串。
	return zapcore.NewTimeEncoder(encoderConfig)
	//后面要用的话直接就给编码器赋值为这个函数的返回值（类型为zapcore.Encoder）
	//encoder := getTimeEncoder()
}

func main() {
	//先初始化logger
	InitLogger()
	defer logger.Sync() //记得关，写到磁盘里，拔环

	/*注册到gin里面去*/
	//先新建一个gin
	r := gin.New()
	//为新建的gin注册2个日志中间件，中间件就是改了gin.DefaultWriter的两个记录日志的默认方法
	//可以用三方库https://github.com/gin-contrib/zap
	//也可以自己建，模仿那两个记录日志的方法，见func main下面两个函数
	r.Use(GinLogger(logger), GinRecovery(logger, true)) //logger在初始化之后不用加数据类型，之前var logger *zap.Logger，此处还初始化了logger
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run(":8080")

}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
