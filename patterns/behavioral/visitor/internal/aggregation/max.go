package aggregation

import "github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"

type MaxVisitor struct {
	results []metrics.Result
}

func NewMaxVisitor() *MaxVisitor {
	return &MaxVisitor{}
}

func (v *MaxVisitor) VisitGauge(metric *metrics.GaugeMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "max",
		Value:     max(sampleValues(metric.Samples())),
		Unit:      metric.Unit(),
	})
}

func (v *MaxVisitor) VisitCounter(metric *metrics.CounterMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "max rate",
		Value:     max(counterRates(metric.Samples())),
		Unit:      rateUnit(metric.Unit()),
	})
}

func (v *MaxVisitor) Results() []metrics.Result {
	return copyResults(v.results)
}
