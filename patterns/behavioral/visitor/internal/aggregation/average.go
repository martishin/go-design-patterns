package aggregation

import "github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"

type AverageVisitor struct {
	results []metrics.Result
}

func NewAverageVisitor() *AverageVisitor {
	return &AverageVisitor{}
}

func (v *AverageVisitor) VisitGauge(metric *metrics.GaugeMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "average",
		Value:     average(sampleValues(metric.Samples())),
		Unit:      metric.Unit(),
	})
}

func (v *AverageVisitor) VisitCounter(metric *metrics.CounterMetric) {
	v.results = append(v.results, metrics.Result{
		Metric:    metric.Name(),
		Operation: "average rate",
		Value:     average(counterRates(metric.Samples())),
		Unit:      rateUnit(metric.Unit()),
	})
}

func (v *AverageVisitor) Results() []metrics.Result {
	return copyResults(v.results)
}
