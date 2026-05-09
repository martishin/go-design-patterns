package aggregation

import "github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"

type MinVisitor struct {
	results []metrics.Result
}

func NewMinVisitor() *MinVisitor {
	return &MinVisitor{}
}

func (v *MinVisitor) VisitGauge(metric *metrics.GaugeMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "min",
		Value:     min(sampleValues(metric.Samples())),
		Unit:      metric.Unit(),
	})
}

func (v *MinVisitor) VisitCounter(metric *metrics.CounterMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "min rate",
		Value:     min(counterRates(metric.Samples())),
		Unit:      rateUnit(metric.Unit()),
	})
}

func (v *MinVisitor) Results() []metrics.Result {
	return copyResults(v.results)
}
