package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
)

var (
	logger Logger

	//zap method
	Binary     = zap.Binary
	Bool       = zap.Bool
	Complex128 = zap.Complex128
	Complex64  = zap.Complex64
	Float64    = zap.Float64
	Float32    = zap.Float32
	Int        = zap.Int
	Int64      = zap.Int64
	Int32      = zap.Int32
	Int16      = zap.Int16
	Int8       = zap.Int8
	String     = zap.String
	Uint       = zap.Uint
	Uint64     = zap.Uint64
	Uint32     = zap.Uint32
	Uint16     = zap.Uint16
	Uint8      = zap.Uint8
	Time       = zap.Time
	Any        = zap.Any
	Duration   = zap.Duration
)

type Logger struct {
	*zap.Logger
}

func Info(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Fatal(msg, fields...)
}

func With(fields ...zap.Field) {
	defer logger.Sync()
	logger.With(fields...)
}

func Sync() {
	logger.Sync()
}

func init() {

	hook := lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     conf.MaxAge,   //days
		Compress:   conf.Compress, // disabled by default
		LocalTime:  true,
	}

	fileWriter := zapcore.AddSync(&hook)

	consoleDebugging := zapcore.Lock(os.Stdout)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	var core zapcore.Core
	if conf.EnableAtomicLevel {
		core = zapcore.NewTee(
			// 打印在控制台
			zapcore.NewCore(encoder, consoleDebugging, level),
			// 打印在文件中
			zapcore.NewCore(encoder, fileWriter, level),
		)
	} else {
		// 仅打印Info级别以上的日志
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		})
		// 打印所有级别的日志
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		})

		core = zapcore.NewTee(
			// 打印在控制台
			zapcore.NewCore(encoder, consoleDebugging, lowPriority),
			// 打印在文件中
			zapcore.NewCore(encoder, fileWriter, highPriority),
		)
	}
	caller := zap.AddCaller()
	callerSkipOpt := zap.AddCallerSkip(1)
	// From a zapcore.Core, it's easy to construct a Logger.
	zapLogger := zap.New(core, caller, callerSkipOpt, zap.AddStacktrace(zap.ErrorLevel))

	logger = Logger{
		zapLogger,
	}

	if conf.EnableAtomicLevel {
		go func() {
			// curl -X PUT -H "Content-Type:application/json" -d '{"level":"info"}' localhost:9090
			http.ListenAndServe(":9090", &level)
		}()
	}
}
