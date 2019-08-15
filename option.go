package trait

import (
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"time"
)

// RotateLogOption ...
type RotateLogOption func(opts *RotateLogOptions)

// ElasticLogOption ...
type ElasticLogOption func(opts *ElasticLogOptions)

func newRotateLogOptions(opts ...RotateLogOption) *RotateLogOptions {
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

// RotateLogLevel ...
func RotateLogLevel(l int) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.Level = l
	}
}

// MaxAge ...
func MaxAge(duration time.Duration) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.MaxAge = duration
	}
}

// RotationTime ...
func RotationTime(duration time.Duration) RotateLogOption {
	return func(opts *RotateLogOptions) {
		opts.RotationTime = duration
	}
}

func newElasticLogOption(opts ...ElasticLogOption) *ElasticLogOptions {
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

// ClientOptions ...
func ClientOptions(cls ...elastic.ClientOptionFunc) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.ClientOptions = cls
	}
}

// Host ...
func Host(h string) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Host = h
	}
}

// Formatter ...
func Formatter(f log.Formatter) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Formatter = f
	}
}

// ReportCaller ...
func ReportCaller(b bool) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.ReportCaller = b
	}
}

// ElasticLogLevel ...
func ElasticLogLevel(l log.Level) ElasticLogOption {
	return func(opts *ElasticLogOptions) {
		opts.Level = l
	}
}
