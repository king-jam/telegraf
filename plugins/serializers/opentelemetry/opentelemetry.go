package opentelemetry

import (
	"github.com/influxdata/telegraf"
	otel "github.com/influxdata/telegraf/plugins/common/opentelemetry/metrics/v1"
)

// Serializer implements the
type Serializer struct {
}

func NewSerializer() (*Serializer, error) {
	s := &Serializer{}
	return s, nil
}

func (s *Serializer) Serialize(metric telegraf.Metric) ([]byte, error) {
	return s.SerializeBatch([]telegraf.Metric{metric})
}

func (s *Serializer) SerializeBatch(metrics []telegraf.Metric) ([]byte, error) {
	//var oMetrics []otel.Metric

	for _, metric := range metrics {
		var oMetric otel.Metric
		oMetric.Name = metric.Name()
		//oMetric.Description	= metric.
		// oMetric.Unit = metric.
		//oMetric.
	}
	return nil, nil
}
