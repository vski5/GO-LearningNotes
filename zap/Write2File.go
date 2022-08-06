package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//用zap.New(…)方法来手动自定义传递所有配置，而不是使用像zap.NewProduction()这样的预置方法来创建logger。
//func New(core zapcore.Core, options ...Option) *Logger
//New接收两个参数，其中有个core zapcore.Core
//zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel。
/*
Encoder:编码器(如何写入日志)。
WriterSyncer ：指定日志将写到哪里去。
Log Level：哪种级别的日志将被写入。
*/

//Encoder:编码器(如何写入日志)
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

//Log Level：哪种级别的日志将被写入。
//本次没用上，直接用官方的日志级别zapcore.DebugLevel
/*
func getLogLevel() zapcore.Level {
	return zapcore.DebugLevel
}
*/
//定义一个全局变量，用来初始化logger，这样就不用每次都初始化了。
var logger *zap.Logger

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
	InitLogger()
	defer logger.Sync() //记得关，写到磁盘里，拔环

	//一个例子：
	eg, err := http.Get("http://www.baidu.com")
	//有错误的情况下，打印错误信息
	if err != nil {
		logger.Error( //记录错误信息,是分级别类型记录
			"Error fetching url",                      //记录的信息
			zap.String("url", "http://www.baidu.com"), //zap.String是一个函数，用来记录字符串，参数是key和value，key是日志的标签，value是日志的内容
			zap.Error(err))                            //zap.Error是一个函数，用来记录错误信息
	} else {
		logger.Info( //记录成功信息
			"Successfully fetched url",
			zap.String("url", "http://www.baidu.com"))
	}
	eg.Body.Close() //关闭body,不然会一直占用内存，body是http.Get求出来的内容。

}
