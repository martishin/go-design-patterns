package aggregation_test

import (
	"math"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/internal/aggregation"
	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

func TestMaxVisitor_VisitsGaugeAndCounter(t *testing.T) {
	gauge := newMaxGaugeMetric(t)
	counter := newMaxCounterMetric(t)
	visitor := aggregation.NewMaxVisitor()

	gauge.Accept(visitor)
	counter.Accept(visitor)

	results := visitor.Results()
	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}

	assertMaxResult(t, results[0], metrics.Result{
		Metric:    "cpu_usage",
		Operation: "max",
		Value:     80,
		Unit:      "percent",
	})
	assertMaxResult(t, results[1], metrics.Result{
		Metric:    "http_requests_total",
		Operation: "max rate",
		Value:     6,
		Unit:      "requests/sec",
	})
}

func newMaxGaugeMetric(t *testing.T) *metrics.GaugeMetric {
	t.Helper()

	gauge, err := metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{
		{Timestamp: 0, Value: 40},
		{Timestamp: 10, Value: 60},
		{Timestamp: 20, Value: 80},
	})
	if err != nil {
		t.Fatalf("NewGaugeMetric() returned error: %v", err)
	}

	return gauge
}

func newMaxCounterMetric(t *testing.T) *metrics.CounterMetric {
	t.Helper()

	counter, err := metrics.NewCounterMetric("http_requests_total", "requests", []metrics.Sample{
		{Timestamp: 0, Value: 100},
		{Timestamp: 10, Value: 130},
		{Timestamp: 20, Value: 190},
	})
	if err != nil {
		t.Fatalf("NewCounterMetric() returned error: %v", err)
	}

	return counter
}

func assertMaxResult(t *testing.T, got metrics.Result, want metrics.Result) {
	t.Helper()

	if got.Metric != want.Metric {
		t.Fatalf("got metric %q, want %q", got.Metric, want.Metric)
	}
	if got.Operation != want.Operation {
		t.Fatalf("got operation %q, want %q", got.Operation, want.Operation)
	}
	if math.Abs(got.Value-want.Value) > 0.0001 {
		t.Fatalf("got value %.2f, want %.2f", got.Value, want.Value)
	}
	if got.Unit != want.Unit {
		t.Fatalf("got unit %q, want %q", got.Unit, want.Unit)
	}
}
