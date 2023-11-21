package log

import (
	"io"
	"os"
	"path"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/log/lreader"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogOptions struct {
	LogdSecret string
	Folder     string
	FilePrefix string
	LogdPort   string
	NodeId     int64
	Rotating   bool
}

var _ logx.Service = (*LogService)(nil)

type LogService struct {
	opts         LogOptions
	appLogger    zerolog.Logger
	engineLogger zerolog.Logger
	siteLogger   zerolog.Logger

	lproxy logx.Proxy
}

func New(opts LogOptions) *LogService {

	if opts.Folder == "" {
		opts.Folder = "tmp/logs"
	}
	if opts.FilePrefix == "" {
		opts.FilePrefix = "temphia_log.log"
	}

	wd, _ := os.Getwd()

	actualPath := path.Join(wd, opts.Folder, opts.FilePrefix)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var logwriter io.WriteCloser
	var lr *lreader.Lreader

	if opts.Rotating {
		file, err := os.OpenFile(actualPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		logwriter = file
		lr = lreader.New(actualPath)
	} else {
		logRotater := &lumberjack.Logger{
			Filename:   actualPath,
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   false,
		}

		logwriter = logRotater
	}

	root := zerolog.New(zerolog.MultiLevelWriter(logwriter, zerolog.NewConsoleWriter())).
		Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
			e.Str("log_event_id", xid.New().String())
			e.Int64("node_id", opts.NodeId)
			e.Timestamp()
		}))

	return &LogService{
		opts:         opts,
		appLogger:    root.With().Str("index", "app").Logger(),
		engineLogger: root.With().Str("index", "engine").Logger(),
		siteLogger:   root.With().Str("index", "site").Logger(),
		lproxy:       lr,
	}
}

func (ls *LogService) GetEngineLogger() *zerolog.Logger {
	return &ls.engineLogger
}

func (ls *LogService) GetAppLogger() *zerolog.Logger { return &ls.appLogger }

func (ls *LogService) GetSiteLogger(tenantId, domain string) zerolog.Logger {
	return ls.siteLogger.With().
		Str("tenant_id", tenantId).
		Str("domain", domain).
		Logger()

}

func (ls *LogService) GetServiceLogger(service string) zerolog.Logger {
	return ls.appLogger.With().Str("service_id", service).Logger()
}

func (ls *LogService) GetLogProxy() logx.Proxy {
	return ls.lproxy
}
