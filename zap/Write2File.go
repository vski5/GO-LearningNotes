package main

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

//用zap.New(…)方法来手动传递所有配置，而不是使用像zap.NewProduction()这样的预置方法来创建logger。
//func New(core zapcore.Core, options ...Option) *Logger
