package metrics

type GaugeMetric struct {
	name    string
	unit    string
	samples []Sample
}

func NewGaugeMetric(name string, unit string, samples []Sample) (*GaugeMetric, error) {
	if err := validateMetric(name, unit, samples); err != nil {
		return nil, err
	}

	return &GaugeMetric{
		name:    name,
		unit:    unit,
		samples: copySamples(samples),
	}, nil
}

func (g *GaugeMetric) Name() string {
	return g.name
}

func (g *GaugeMetric) Unit() string {
	return g.unit
}

func (g *GaugeMetric) Samples() []Sample {
	return copySamples(g.samples)
}

func (g *GaugeMetric) Accept(visitor MetricVisitor) {
	visitor.VisitGauge(g)
}
