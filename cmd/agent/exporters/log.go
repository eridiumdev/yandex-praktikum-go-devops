package exporters

import (
	"eridiumdev/yandex-praktikum-go-devops/internal/logger"
	"eridiumdev/yandex-praktikum-go-devops/internal/metrics"
)

type LogExporter struct {
	*AbstractExporter
}

func NewLogExporter(name string) *LogExporter {
	exp := &LogExporter{
		AbstractExporter: &AbstractExporter{
			name:  name,
			ready: make(chan bool),
		},
	}
	exp.makeReady()
	return exp
}

func (exp *LogExporter) Export(mtx []metrics.Metric) error {
	defer func() {
		exp.makeReady()
	}()

	for _, metric := range mtx {
		logger.Infof("%s:%s (%s)", metric.GetName(), metric.GetStringValue(), metric.GetType())
	}
	return nil
}