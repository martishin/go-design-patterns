package aggregation_test

import (
	"math"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/internal/aggregation"
	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

func TestMinVisitor_VisitsGaugeAndCounter(t *testing.T) {
	gauge := newMinGaugeMetric(t)
	counter := newMinCounterMetric(t)
	visitor := aggregation.NewMinVisitor()

	gauge.Accept(visitor)
	counter.Accept(visitor)

	results := visitor.Results()
	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}

	assertMinResult(t, results[0], metrics.Result{
		Metric:    "cpu_usage",
		Operation: "min",
		Value:     40,
		Unit:      "percent",
	})
	assertMinResult(t, results[1], metrics.Result{
		Metric:    "http_requests_total",
		Operation: "min rate",
		Value:     3,
		Unit:      "requests/sec",
	})
}

func newMinGaugeMetric(t *testing.T) *metrics.GaugeMetric {
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

func newMinCounterMetric(t *testing.T) *metrics.CounterMetric {
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

func assertMinResult(t *testing.T, got metrics.Result, want metrics.Result) {
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
