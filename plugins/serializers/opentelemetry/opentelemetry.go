package opentelemetry

import "github.com/influxdata/telegraf"

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
	return nil, nil
}
