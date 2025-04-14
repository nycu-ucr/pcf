package app

import (
	pcf_context "github.com/nycu-ucr/pcf/internal/context"
	"github.com/nycu-ucr/pcf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *pcf_context.PCFContext
	Config() *factory.Config
}
