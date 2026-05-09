package aggregation_test

import (
	"math"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/internal/aggregation"
	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

func TestAverageVisitor_VisitsGaugeAndCounter(t *testing.T) {
	gauge := newAverageGaugeMetric(t)
	counter := newAverageCounterMetric(t)
	visitor := aggregation.NewAverageVisitor()

	gauge.Accept(visitor)
	counter.Accept(visitor)

	results := visitor.Results()
	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}

	assertAverageResult(t, results[0], metrics.Result{
		Metric:    "cpu_usage",
		Operation: "average",
		Value:     60,
		Unit:      "percent",
	})
	assertAverageResult(t, results[1], metrics.Result{
		Metric:    "http_requests_total",
		Operation: "average rate",
		Value:     4.5,
		Unit:      "requests/sec",
	})
}

func TestAverageVisitor_ResultsReturnsCopy(t *testing.T) {
	gauge := newAverageGaugeMetric(t)
	visitor := aggregation.NewAverageVisitor()
	gauge.Accept(visitor)

	results := visitor.Results()
	results[0].Value = 99

	if visitor.Results()[0].Value != 60 {
		t.Fatalf("got result value %.2f, want 60.00", visitor.Results()[0].Value)
	}
}

func newAverageGaugeMetric(t *testing.T) *metrics.GaugeMetric {
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

func newAverageCounterMetric(t *testing.T) *metrics.CounterMetric {
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

func assertAverageResult(t *testing.T, got metrics.Result, want metrics.Result) {
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
