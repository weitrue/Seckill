package xzap

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/weitrue/Seckill/pkg/logger/meta"
)

const (
	debugFilename  = "debug.log"
	accessFilename = "access.log"
	errorFilename  = "error.log"
	severeFilename = "severe.log"

	consoleMode = "console"
	volumeMode  = "volume"

	maxSize   = 30
	maxBackup = 5

	DefaultTimeLayout = "2006-01-02T15:04:05.000Z07"
	DefaultKeepDays   = 7
)

var (
	// ErrLogPathNotSet is an error that indicates the log path is not set.
	ErrLogPathNotSet = errors.New("log path must be set")
	// ErrLogServiceNameNotSet is an error that indicates that the service name is not set.
	ErrLogServiceNameNotSet = errors.New("log service name must be set")

	Logger *ZapLogger

	once sync.Once
)

// ZapLogger zap日志
type ZapLogger struct {
	logger *zap.Logger
	Opt    *Opt
}

// SetUp 初始化zap Logger
func SetUp(c Config, opts ...Option) (*ZapLogger, error) {
	if c.KeepDays == 0 {
		c.KeepDays = DefaultKeepDays
	}

	opt := &Opt{
		fields: make(map[string]string),
	}

	for _, f := range opts {
		if f != nil {
			f(opt)
		}
	}

	if len(c.Path) == 0 {
		return nil, ErrLogPathNotSet
	}

	if len(c.ServiceName) == 0 {
		return nil, ErrLogServiceNameNotSet
	}

	switch c.Mode {
	case consoleMode:
		withLogLevel(c.Level)
		setupWithConsole(opt)
	case volumeMode:
		setupWithFiles(c, opt)
	default:
		setupWithFiles(c, opt)
	}

	return Logger, nil
}

func (zl *ZapLogger) GetLogger() *zap.Logger {
	return zl.logger
}

func (zl *ZapLogger) With(data ...meta.Field) *ZapLogger {
	zl.logger = zl.logger.With(wrapZapMeta(nil, data...)...)
	return zl
}

// Print debug log
func (zl *ZapLogger) Print(msg string, level zapcore.Level, err error, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Check(level, msg).Write(wrapZapMeta(err, data...)...)
}

// Debug debug log
func (zl *ZapLogger) Debug(msg string, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, wrapZapMeta(nil, data...)...)
}

// Info log
func (zl *ZapLogger) Info(msg string, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, wrapZapMeta(nil, data...)...)
}

// Warn log
func (zl *ZapLogger) Warn(msg string, err error, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, wrapZapMeta(err, data...)...)
}

// Error error log
func (zl *ZapLogger) Error(msg string, err error, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, wrapZapMeta(err, data...)...)
}

// Panic panic log
func (zl *ZapLogger) Panic(msg string, err error, data ...meta.Field) {
	zl.logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, wrapZapMeta(err, data...)...)
}

// Debugf debug
func (zl *ZapLogger) Debugf(msg, format string, data ...any) {
	zl.Debug(msg, nil, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Infof info
func (zl *ZapLogger) Infof(msg, format string, data ...any) {
	zl.Info(msg, nil, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Warnf warn
func (zl *ZapLogger) Warnf(msg, format string, err error, data ...any) {
	zl.Warn(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Errorf error
func (zl *ZapLogger) Errorf(msg, format string, err error, data ...any) {
	zl.Error(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Panicf panic
func (zl *ZapLogger) Panicf(msg, format string, err error, data ...any) {
	zl.Panic(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// wrapZapMeta wrap meta to zap fields
func wrapZapMeta(err error, metas ...meta.Field) (fields []zap.Field) {
	capacity := len(metas) + 1 // namespace meta
	if err != nil {
		capacity++
	}

	fields = make([]zap.Field, 0, capacity)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	for _, m := range metas {
		fields = append(fields, zap.Any(m.Key(), m.Value()))
	}

	return
}

func setupWithConsole(opt *Opt) {
	consoleDebugging := zapcore.Lock(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(getConsoleEncoder(), consoleDebugging, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= opt.level
		})),
	)

	once.Do(func() {
		log := zap.New(core, zap.AddCaller())
		for key, value := range opt.fields {
			log = log.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
		}

		Logger = &ZapLogger{
			logger: log,
			Opt:    opt,
		}
	})

}

func setupWithFiles(c Config, opt *Opt) {
	accessPath := path.Join(c.Path, accessFilename)
	errorPath := path.Join(c.Path, errorFilename)
	severePath := path.Join(c.Path, severeFilename)
	debugPath := path.Join(c.Path, debugFilename)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(getFileEncoder(), getLogWriter(accessPath, maxSize, maxBackup, c.KeepDays, c.Compress), infoPriority),
		zapcore.NewCore(getFileEncoder(), getLogWriter(errorPath, maxSize, maxBackup, c.KeepDays, c.Compress), errPriority),
		zapcore.NewCore(getFileEncoder(), getLogWriter(severePath, maxSize, maxBackup, c.KeepDays, c.Compress), warnPriority),
		zapcore.NewCore(getFileEncoder(), getLogWriter(debugPath, maxSize, maxBackup, c.KeepDays, c.Compress), debugPriority),
	)

	stderr := zapcore.Lock(os.Stderr) // lock for concurrent safe

	once.Do(func() {
		log := zap.New(core, zap.AddCaller(), zap.ErrorOutput(stderr))
		for key, value := range opt.fields {
			log = log.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
		}

		Logger = &ZapLogger{
			logger: log,
			Opt:    opt,
		}
	})
}

func getLogWriter(fileName string, maxSize, maxBackups, maxAge int, isCompress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   isCompress,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getFileEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "Logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // Level 序列化为小写字符串
		EncodeTime:     TimeEncoder,                    // 记录时间设置为2006-01-02T15:04:05Z07:00
		EncodeDuration: zapcore.SecondsDurationEncoder, //  耗时设置为浮点秒数
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	})
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// TimeEncoder 设置时间格式化方式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(DefaultTimeLayout))
}
