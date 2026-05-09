package metrics

import "errors"

var (
	ErrEmptyMetricName = errors.New("metrics: name is empty")
	ErrEmptyMetricUnit = errors.New("metrics: unit is empty")
	ErrNoMetricSamples = errors.New("metrics: samples are empty")
)

type Sample struct {
	Timestamp int
	Value     float64
}

type Result struct {
	Metric    string
	Operation string
	Value     float64
	Unit      string
}

type Metric interface {
	Name() string
	Unit() string
	Samples() []Sample
	Accept(MetricVisitor)
}

type MetricVisitor interface {
	VisitGauge(*GaugeMetric)
	VisitCounter(*CounterMetric)
}

func validateMetric(name string, unit string, samples []Sample) error {
	if name == "" {
		return ErrEmptyMetricName
	}
	if unit == "" {
		return ErrEmptyMetricUnit
	}
	if len(samples) == 0 {
		return ErrNoMetricSamples
	}

	return nil
}

func copySamples(samples []Sample) []Sample {
	return append([]Sample(nil), samples...)
}
