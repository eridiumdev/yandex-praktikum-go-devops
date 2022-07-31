package exporters

import (
	"context"

	"eridiumdev/yandex-praktikum-go-devops/internal/common/executor"
	"eridiumdev/yandex-praktikum-go-devops/internal/common/logger"
	"eridiumdev/yandex-praktikum-go-devops/internal/metrics/domain"
)

type LogExporter struct {
	*executor.Executor
}

func NewLogExporter(name string) *LogExporter {
	exp := &LogExporter{
		Executor: executor.New(name),
	}
	exp.ReadyUp()
	return exp
}

func (exp *LogExporter) Export(ctx context.Context, mtx []domain.Metric) error {
	defer func() {
		exp.ReadyUp()
	}()

	for _, metric := range mtx {
		logger.New(ctx).Infof("%s:%s (%s)", metric.Name, metric.StringValue(), metric.Type)
	}
	return nil
}