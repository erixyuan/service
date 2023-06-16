package global

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"service/interfaces"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
	Config    *LoggerConfig
}

var instance *Logger

func GetLogger() *Logger {
	if instance == nil {
		instance = &Logger{}
		instance.initLogCompnent()
	}
	return instance
}

func InitLoginConfig(maxAge int) *Logger {
	fmt.Printf("初始化日志配置 maxAge:%d", maxAge)
	instance = &Logger{}
	instance.Config = &LoggerConfig{}
	instance.Config.MaxAge = maxAge
	instance.initLogCompnent()
	return instance
}

func (logger *Logger) Info(msg string) {
	logger.zapLogger.Info(msg)
}

func (logger *Logger) InfoByTrace(base interfaces.BaseInterface, msg string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	logger.zapLogger.Info(
		fmt.Sprintf(msg, args...),
		zap.String("log_type", "process"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", base.GetTraceId()),
		zap.String("request_id", base.GetRequestId()),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) Infof(msg string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	ctx := context.Background()
	var traceId string
	var requestId string
	if ok := ctx.Value("trace_id"); ok != nil {
		traceId = ok.(string)
	}
	if ok := ctx.Value("request_id"); ok != nil {
		requestId = ok.(string)
	}
	logger.zapLogger.Info(
		fmt.Sprintf(msg, args...),
		zap.String("log_type", "process"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", traceId),
		zap.String("request_id", requestId),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) ErrorByTrace(base interfaces.BaseInterface, msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.zapLogger.Error(
		fmt.Sprintf(msg, args...),
		zap.String("log_type", "process"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", base.GetTraceId()),
		zap.String("request_id", base.GetRequestId()),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) Error(msg string) {
	logger.zapLogger.Error(msg)
}

func (logger *Logger) Errorf(msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	ctx := context.Background()
	var traceId string
	var requestId string
	if ok := ctx.Value("trace_id"); ok != nil {
		traceId = ok.(string)
	}
	if ok := ctx.Value("request_id"); ok != nil {
		requestId = ok.(string)
	}
	logger.zapLogger.Error(
		fmt.Sprintf(msg, args...),
		zap.String("log_type", "process"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", traceId),
		zap.String("request_id", requestId),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) WarnByTrace(base interfaces.BaseInterface, msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.zapLogger.Warn(
		fmt.Sprintf(msg, args...),
		zap.String("log_type", "process"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", base.GetTraceId()),
		zap.String("request_id", base.GetRequestId()),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) InfoRequest(msg, traceId, requestId, method, ip, url, reqBody string, statusCode int) {
	_, file, line, _ := runtime.Caller(1)
	logger.zapLogger.Info(
		msg+url+"|"+reqBody,
		zap.String("log_type", "in"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", traceId),
		zap.String("request_id", requestId),
		zap.Int64("latency", 0),
	)
}

func (logger *Logger) InfoResponse(msg, traceId, requestId, method, ip, url, reqBody, respBody string, statusCode int, latencyTime int64) {
	_, file, line, _ := runtime.Caller(1)
	logger.zapLogger.Info(
		msg+url+"|"+respBody,
		zap.String("log_type", "out"),
		zap.String("line", file+" "+strconv.Itoa(line)),
		zap.String("trace_id", traceId),
		zap.String("request_id", requestId),
		zap.Int64("latency", latencyTime),
	)
}

func (logger *Logger) getWriter(filename string) io.Writer {
	maxAge := 30 // 默认保存30天
	if logger.Config != nil {
		maxAge = logger.Config.MaxAge
	}

	// filename是指向最新日志的链接
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxAge)), // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24),                 //切割频率 24小时
	)
	if err != nil {
		log.Println("Logger Start Error")
		panic(err)
	}
	return hook
}

func (logger *Logger) initLogger(allLogPath, infoLogPath, errLogPath string, logLevel zapcore.Level) *zap.Logger {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		CallerKey:    "file",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, //输出的时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	allWriter := logger.getWriter(allLogPath)
	infoWriter := logger.getWriter(infoLogPath)
	warnWriter := logger.getWriter(errLogPath)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(allWriter), allLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func (logger *Logger) initLogCompnent() {
	workDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logsAllPath := path.Join(workDir, "logs/all.log")
	logsInfoPath := path.Join(workDir, "logs/info.log")
	logsErrorPath := path.Join(workDir, "logs/error.log")
	logger.zapLogger = logger.initLogger(logsAllPath, logsInfoPath, logsErrorPath, zap.DebugLevel)
	defer logger.zapLogger.Sync()
	logger.zapLogger.Sugar().Infof("Log init success")
}
