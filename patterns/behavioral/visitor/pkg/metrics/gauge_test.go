package metrics_test

import (
	"errors"
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

type gaugeVisitorSpy struct {
	visited *metrics.GaugeMetric
}

func (v *gaugeVisitorSpy) VisitGauge(metric *metrics.GaugeMetric) {
	v.visited = metric
}

func (v *gaugeVisitorSpy) VisitCounter(_ *metrics.CounterMetric) {}

func TestNewGaugeMetric(t *testing.T) {
	gauge, err := metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{
		{Timestamp: 0, Value: 40},
		{Timestamp: 10, Value: 60},
	})
	if err != nil {
		t.Fatalf("NewGaugeMetric() returned error: %v", err)
	}

	if gauge.Name() != "cpu_usage" {
		t.Fatalf("got name %q, want %q", gauge.Name(), "cpu_usage")
	}
	if gauge.Unit() != "percent" {
		t.Fatalf("got unit %q, want %q", gauge.Unit(), "percent")
	}
	if len(gauge.Samples()) != 2 {
		t.Fatalf("got %d samples, want 2", len(gauge.Samples()))
	}
}

func TestNewGaugeMetric_ValidatesInput(t *testing.T) {
	tests := []struct {
		name    string
		metric  string
		unit    string
		samples []metrics.Sample
		wantErr error
	}{
		{
			name:    "empty name",
			unit:    "percent",
			samples: []metrics.Sample{{Timestamp: 0, Value: 40}},
			wantErr: metrics.ErrEmptyMetricName,
		},
		{
			name:    "empty unit",
			metric:  "cpu_usage",
			samples: []metrics.Sample{{Timestamp: 0, Value: 40}},
			wantErr: metrics.ErrEmptyMetricUnit,
		},
		{
			name:    "empty samples",
			metric:  "cpu_usage",
			unit:    "percent",
			wantErr: metrics.ErrNoMetricSamples,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := metrics.NewGaugeMetric(tt.metric, tt.unit, tt.samples)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestGaugeMetric_AcceptCallsVisitGauge(t *testing.T) {
	gauge, err := metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{{Timestamp: 0, Value: 40}})
	if err != nil {
		t.Fatalf("NewGaugeMetric() returned error: %v", err)
	}

	visitor := &gaugeVisitorSpy{}
	gauge.Accept(visitor)

	if visitor.visited != gauge {
		t.Fatal("expected visitor to receive gauge metric")
	}
}

func TestGaugeMetric_SamplesReturnsCopy(t *testing.T) {
	gauge, err := metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{{Timestamp: 0, Value: 40}})
	if err != nil {
		t.Fatalf("NewGaugeMetric() returned error: %v", err)
	}

	samples := gauge.Samples()
	samples[0].Value = 99

	if gauge.Samples()[0].Value != 40 {
		t.Fatalf("got sample value %.2f, want 40.00", gauge.Samples()[0].Value)
	}
}
