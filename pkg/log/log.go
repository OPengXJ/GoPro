package log

import (
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

type Level = zapcore.Level

const (
    InfoLevel   Level = zap.InfoLevel   // 0, default level
    WarnLevel   Level = zap.WarnLevel   // 1
    ErrorLevel  Level = zap.ErrorLevel  // 2
    DPanicLevel Level = zap.DPanicLevel // 3, used in development log
    // PanicLevel logs a message, then panics
    PanicLevel Level = zap.PanicLevel // 4
    // FatalLevel logs a message, then calls os.Exit(1).
    FatalLevel Level = zap.FatalLevel // 5
    DebugLevel Level = zap.DebugLevel // -1
)

var (
    Skip        = zap.Skip
    Binary      = zap.Binary
    Bool        = zap.Bool
    Boolp       = zap.Boolp
    ByteString  = zap.ByteString
	String		= zap.String
	Strings		= zap.Strings

    Float64     = zap.Float64
    Float64p    = zap.Float64p
    Float32     = zap.Float32
    Float32p    = zap.Float32p
    Durationp   = zap.Durationp
    Any         = zap.Any

    Info   = log.Info
    Warn   = log.Warn
    Error  = log.Error
    DPanic = log.DPanic
    Panic  = log.Panic
    Fatal  = log.Fatal
    Debug  = log.Debug
)



var log=New()

func New()*zap.Logger{
    var coreArr []zapcore.Core

    //获取编码器
    encoderConfig := zap.NewProductionEncoderConfig()               //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder           //指定时间格式
    encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder    //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
    //encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
    encoder := zapcore.NewConsoleEncoder(encoderConfig)

    //日志级别
    highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool{  //error级别
        return lev >= zap.ErrorLevel
    })
    lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {  //info和debug级别,debug级别是最低的
        return lev < zap.ErrorLevel && lev >= zap.DebugLevel
    })
	zap.Skip()
    //info文件writeSyncer
    infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
        Filename:   "./pkg/log/info.log",   //日志文件存放目录，如果文件夹不存在会自动创建
        MaxSize:    2,                  //文件大小限制,单位MB
        MaxBackups: 100,                //最大保留日志文件数量
        MaxAge:     30,                 //日志文件保留天数
        Compress:   false,              //是否压缩处理
    })
    infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer,zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
    //error文件writeSyncer
    errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
        Filename:   "./pkg/log/error.log",      //日志文件存放目录
        MaxSize:    1,                      //文件大小限制,单位MB
        MaxBackups: 5,                      //最大保留日志文件数量
        MaxAge:     30,                     //日志文件保留天数
        Compress:   false,                  //是否压缩处理
    })
    errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer,zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

    coreArr = append(coreArr, infoFileCore)
    coreArr = append(coreArr, errorFileCore)
    log := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
	return log
}