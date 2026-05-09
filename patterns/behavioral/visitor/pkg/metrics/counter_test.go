package metrics_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

type counterVisitorSpy struct {
	visited *metrics.CounterMetric
}

func (v *counterVisitorSpy) VisitGauge(_ *metrics.GaugeMetric) {}

func (v *counterVisitorSpy) VisitCounter(metric *metrics.CounterMetric) {
	v.visited = metric
}

func TestNewCounterMetric(t *testing.T) {
	counter, err := metrics.NewCounterMetric("http_requests_total", "requests", []metrics.Sample{
		{Timestamp: 0, Value: 100},
		{Timestamp: 10, Value: 130},
	})
	if err != nil {
		t.Fatalf("NewCounterMetric() returned error: %v", err)
	}

	if counter.Name() != "http_requests_total" {
		t.Fatalf("got name %q, want %q", counter.Name(), "http_requests_total")
	}
	if counter.Unit() != "requests" {
		t.Fatalf("got unit %q, want %q", counter.Unit(), "requests")
	}
	if len(counter.Samples()) != 2 {
		t.Fatalf("got %d samples, want 2", len(counter.Samples()))
	}
}

func TestNewCounterMetric_ValidatesInput(t *testing.T) {
	tests := []struct {
		name    string
		samples []metrics.Sample
		wantErr error
	}{
		{
			name:    "one sample",
			samples: []metrics.Sample{{Timestamp: 0, Value: 100}},
			wantErr: metrics.ErrCounterNeedsTwoSamples,
		},
		{
			name: "timestamps do not increase",
			samples: []metrics.Sample{
				{Timestamp: 10, Value: 100},
				{Timestamp: 10, Value: 130},
			},
			wantErr: metrics.ErrCounterSampleOrder,
		},
		{
			name: "counter decreases",
			samples: []metrics.Sample{
				{Timestamp: 0, Value: 100},
				{Timestamp: 10, Value: 90},
			},
			wantErr: metrics.ErrCounterDecrease,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := metrics.NewCounterMetric("http_requests_total", "requests", tt.samples)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestCounterMetric_AcceptCallsVisitCounter(t *testing.T) {
	counter, err := metrics.NewCounterMetric(
		"http_requests_total",
		"requests",
		[]metrics.Sample{{Timestamp: 0, Value: 100}, {Timestamp: 10, Value: 130}},
	)
	if err != nil {
		t.Fatalf("NewCounterMetric() returned error: %v", err)
	}

	visitor := &counterVisitorSpy{}
	counter.Accept(visitor)

	if visitor.visited != counter {
		t.Fatal("expected visitor to receive counter metric")
	}
}

func TestCounterMetric_SamplesReturnsCopy(t *testing.T) {
	counter, err := metrics.NewCounterMetric(
		"http_requests_total",
		"requests",
		[]metrics.Sample{{Timestamp: 0, Value: 100}, {Timestamp: 10, Value: 130}},
	)
	if err != nil {
		t.Fatalf("NewCounterMetric() returned error: %v", err)
	}

	samples := counter.Samples()
	samples[0].Value = 999

	if counter.Samples()[0].Value != 100 {
		t.Fatalf("got sample value %.2f, want 100.00", counter.Samples()[0].Value)
	}
}
