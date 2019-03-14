package trait

import (
	"fmt"
	"github.com/godcong/elogrus"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

type ElasticLogOptions struct {
	ReportCaller  bool
	Formatter     log.Formatter
	Host          string
	Level         log.Level
	ClientOptions []elastic.ClientOptionFunc
}

func InitElasticLog(index string, opts ...ElasticLogOption) {
	opt := newElasticLogOption(opts)
	client, err := elastic.NewClient(opt.ClientOptions...)
	if err != nil {
		log.Panic(err)
	}

	t, err := elogrus.NewElasticHook(client, opt.Host, opt.Level, index)
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(t)

	log.SetReportCaller(opt.ReportCaller)
	log.SetFormatter(opt.Formatter)
}

type RotateLogOptions struct {
	Level        int           `json:"level"`
	MaxAge       time.Duration `json:"max_age"`
	RotationTime time.Duration `json:"rotation_time"`
}

const (
	RotateLogAll = iota
	RotateLogTrace
	RotateLogDebug
	RotateLogInfo
	RotateLogWarn
	RotateLogError
	RotateLogFatal
	RotateLogPanic
	RotateLogOff
)

// InitRotateLogger ...
func InitRotateLog(logPath string, opts ...RotateLogOption) {
	opt := newRotateLogOptions(opts)
	dir, filename := filepath.Split(logPath)
	_ = os.MkdirAll(dir, os.ModePerm)
	writer, err := rotatelogs.New(
		dir+"%Y%m%d%H%M_"+filename,
		rotatelogs.WithLinkName(logPath),              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(opt.MaxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(opt.RotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	switch opt.Level {
	case RotateLogDebug:
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stderr)
	case RotateLogInfo:
		NoOutput()
		log.SetLevel(log.InfoLevel)
	case RotateLogWarn:
		NoOutput()
		log.SetLevel(log.WarnLevel)
	case RotateLogError:
		NoOutput()
		log.SetLevel(log.ErrorLevel)
	default:
		NoOutput()
		log.SetLevel(log.TraceLevel)
	}

	//hook := lfshook.NewHook(lfshook.WriterMap{
	//	log.DebugLevel: writer, // 为不同级别设置不同的输出目的
	//	log.InfoLevel:  writer,
	//	log.WarnLevel:  writer,
	//	log.ErrorLevel: writer,
	//	log.FatalLevel: writer,
	//	log.PanicLevel: writer,
	//}, &log.JSONFormatter{})
	hook := lfshook.NewHook(writer, &log.JSONFormatter{})
	//hook.SetDefaultWriter(writer)
	//hook.SetFormatter(&log.JSONFormatter{})
	log.AddHook(hook)

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
}

func NoOutput() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	log.SetOutput(src)
}
