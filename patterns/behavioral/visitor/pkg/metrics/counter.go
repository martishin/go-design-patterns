package metrics

import "errors"

var (
	ErrCounterNeedsTwoSamples = errors.New("metrics: counter requires at least two samples")
	ErrCounterSampleOrder     = errors.New("metrics: counter samples must have increasing timestamps")
	ErrCounterDecrease        = errors.New("metrics: counter samples must not decrease")
)

type CounterMetric struct {
	name    string
	unit    string
	samples []Sample
}

func NewCounterMetric(name string, unit string, samples []Sample) (*CounterMetric, error) {
	if err := validateMetric(name, unit, samples); err != nil {
		return nil, err
	}
	if len(samples) < 2 {
		return nil, ErrCounterNeedsTwoSamples
	}
	if err := validateCounterSamples(samples); err != nil {
		return nil, err
	}

	return &CounterMetric{
		name:    name,
		unit:    unit,
		samples: copySamples(samples),
	}, nil
}

func (c *CounterMetric) Name() string {
	return c.name
}

func (c *CounterMetric) Unit() string {
	return c.unit
}

func (c *CounterMetric) Samples() []Sample {
	return copySamples(c.samples)
}

func (c *CounterMetric) Accept(visitor MetricVisitor) {
	visitor.VisitCounter(c)
}

func validateCounterSamples(samples []Sample) error {
	for index := 1; index < len(samples); index++ {
		previous := samples[index-1]
		current := samples[index]

		if current.Timestamp <= previous.Timestamp {
			return ErrCounterSampleOrder
		}
		if current.Value < previous.Value {
			return ErrCounterDecrease
		}
	}

	return nil
}
