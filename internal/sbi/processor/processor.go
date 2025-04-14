package processor

import (
	"github.com/nycu-ucr/pcf/internal/sbi/consumer"
	"github.com/nycu-ucr/pcf/pkg/app"
)

type PCF interface {
	app.App
	Consumer() *consumer.Consumer
}

type Processor struct {
	PCF
}

func NewProcessor(pcf PCF) (*Processor, error) {
	p := &Processor{
		PCF: pcf,
	}

	return p, nil
}
