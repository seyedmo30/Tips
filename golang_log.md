https://www.kelche.co/blog/go/zap/
      
      
      package logger
      
      import (
      	"go.uber.org/zap/zapcore"
      	"os"
      	"sync"
      	"time"
      
      	"go.uber.org/zap"
      )
      
      var once sync.Once
      
      var instance *zap.Logger
      
      func Connect() *zap.Logger {
      
      	once.Do(func() {
      
      		config := zap.NewProductionEncoderConfig()
      		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
      		config.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
      		consoleEncoder := zapcore.NewConsoleEncoder(config)
      		defaultLogLevel := zapcore.InfoLevel
      
      		deepLevel := zapcore.FatalLevel
      		core := zapcore.NewTee(
      			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
      		)
      
      		instance = zap.New(core, zap.AddCaller(), zap.AddStacktrace(deepLevel))
      
      		//logger, _ := zap.NewProduction()
      		//defer logger.Sync() // flushes buffer, if any
      		//logger.Sugar()
      		//instance1 := logger.Sugar()
      		//instance1.Info()
      		//var filename string
      		//err := os.MkdirAll("/var/logger/exploit/", 0555)
      		//if err != nil {
      		//	fmt.Println("MkdirAll : ", err)
      		//	filename = "./logs.logger"
      		//
      		//} else {
      		//	filename = "/var/logger/exploit/logs.logger"
      		//}
      		//filename = "/var/logger/exploit/logs.logger"
      		//
      		//config := zap.NewProductionEncoderConfig()
      		//config.EncodeTime = zapcore.ISO8601TimeEncoder
      		//fileEncoder := zapcore.NewJSONEncoder(config)
      		//logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
      		//if err != nil {
      		//	fmt.Println("os.OpenFile var/logger/exploit/logs.logger ", err)
      		//	filename = "./logs.logger"
      		//
      		//	logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
      		//	if err != nil {
      		//		fmt.Println("os.OpenFile logs.logger ", err)
      		//
      		//		os.Exit(1)
      		//	}
      		//}
      		//writer := zapcore.AddSync(logFile)
      		//defaultLogLevel := zapcore.DebugLevel
      		//core := zapcore.NewTee(
      		//	zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
      		//)
      		//instance = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
      
      	})
      
      	return instance
      }
