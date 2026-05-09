package metrics_test

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

type metricVisitorStub struct {
	gaugeVisited   bool
	counterVisited bool
}

func (v *metricVisitorStub) VisitGauge(_ *metrics.GaugeMetric) {
	v.gaugeVisited = true
}

func (v *metricVisitorStub) VisitCounter(_ *metrics.CounterMetric) {
	v.counterVisited = true
}

func TestMetricVisitor_InterfaceCanBeImplemented(t *testing.T) {
	var visitor metrics.MetricVisitor = &metricVisitorStub{}

	gauge, err := metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{{Timestamp: 0, Value: 50}})
	if err != nil {
		t.Fatalf("NewGaugeMetric() returned error: %v", err)
	}

	counter, err := metrics.NewCounterMetric(
		"http_requests_total",
		"requests",
		[]metrics.Sample{{Timestamp: 0, Value: 100}, {Timestamp: 10, Value: 130}},
	)
	if err != nil {
		t.Fatalf("NewCounterMetric() returned error: %v", err)
	}

	gauge.Accept(visitor)
	counter.Accept(visitor)

	stub := visitor.(*metricVisitorStub)
	if !stub.gaugeVisited {
		t.Fatal("expected gauge visit")
	}
	if !stub.counterVisited {
		t.Fatal("expected counter visit")
	}
}
