package trait

import (
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"time"
)

type RotateLogOption func(opts *RotateLogOptions)
type ElasticLogOption func(opts *ElasticLogOptions)

func newRotateLogOptions(opts []RotateLogOption) *RotateLogOptions {
	opt := &RotateLogOptions{
		Level:        RotateLogAll,
		MaxAge:       30 * 24 * time.Hour, //log saved one month
		RotationTime: 24 * time.Hour,
	}

	for _, o := range opts {
		o(opt)
	}
	return opt
}

func RotateLogLevel(l int) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.Level = l
	}
}

func MaxAge(duration time.Duration) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.MaxAge = duration
	}
}

func RotationTime(duration time.Duration) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.RotationTime = duration
	}
}

func newElasticLogOption(opts []ElasticLogOption) *ElasticLogOptions {
	opt := &ElasticLogOptions{
		ReportCaller: true,
		Formatter:    &log.JSONFormatter{},
		Host:         "localhost",
		Level:        log.TraceLevel,
		ClientOptions: []elastic.ClientOptionFunc{
			elastic.SetSniff(false),
			elastic.SetURL("http://localhost:9200"),
		},
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

func ClientOptions(cls ...elastic.ClientOptionFunc) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.ClientOptions = cls
	}
}

func Host(h string) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Host = h
	}
}

func Formatter(f log.Formatter) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Formatter = f
	}
}

func ReportCaller(b bool) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.ReportCaller = b
	}
}

func ElasticLogLevel(l log.Level) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Level = l
	}
}
