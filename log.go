package trait

import (
	"github.com/godcong/elogrus"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

type ElasticLogOption struct {
	ReportCaller  bool
	Formatter     log.Formatter
	Host          string
	Level         log.Level
	ClientOptions []elastic.ClientOptionFunc
}

func DefaultElasticLogOption(opt *ElasticLogOption) *ElasticLogOption {
	if opt == nil {
		opt = &ElasticLogOption{
			ReportCaller: true,
			Formatter:    &log.JSONFormatter{},
			Host:         "localhost",
			Level:        log.TraceLevel,
			ClientOptions: []elastic.ClientOptionFunc{
				elastic.SetSniff(false),
				elastic.SetURL("http://localhost:9200"),
			},
		}
	}
	return opt
}

func InitElasticLog(index string, opt *ElasticLogOption) {
	opt = DefaultElasticLogOption(opt)
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
