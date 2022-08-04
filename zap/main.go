package main

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

//一个日志库，来自UBer
//三个要求，能记录，可切割分大小，支持不同类别日志级别
/*
开始
*/
//设置输出位置，文件属性
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Ubuntu/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//log.SetOutput()正式的把文件（路径） 设置为日志
	log.SetOutput(logFileLocation) //把日志文件地址传进去
}

//使用日志
//用log.Printf记录日志 使用%s来记录字符串，%d来记录整数，%f来记录浮点数
//log.Printf("Error fetching url %s : %s " , url, err.Error())

//下载 go get

//提供两种类型的日志：Sugared Logger（更快）和Logger

//3种创建logger的方法
//zap.NewProduction()/zap.NewDevelopment()/zap.Example()

//先定义logger全局实例
var logger *zap.Logger

/*Sugared Logger定义方法：
var sugarLogger *zap.SugaredLogger
其他都一样，两个logger都打印输出JSON结构格式的日志
*/

//设置logger的初始化
func InitLogger() {
	logger, _ = zap.NewProduction()
	//对于SugaredLogger还有一步：
	//SugarLogger = Logger.Sugar() 被称为加点糖
	//后面直接用SugarLogger.Infof("Error", zap.String("url", "http://www.baidu.com"), zap.Error(err) ) 之类的就可以了
}

//首先创建一个Logger，然后使用Info/ Error等Logger方法记录消息。
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
	//现在只能从控制台打出来
	//打到日志文件里面去：

}
