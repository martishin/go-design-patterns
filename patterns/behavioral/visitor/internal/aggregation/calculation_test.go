package aggregation

import (
	"testing"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

func TestSampleValues(t *testing.T) {
	values := sampleValues([]metrics.Sample{
		{Timestamp: 0, Value: 40},
		{Timestamp: 10, Value: 60},
	})

	if len(values) != 2 {
		t.Fatalf("got %d values, want 2", len(values))
	}
	if values[0] != 40 || values[1] != 60 {
		t.Fatalf("got values %v, want [40 60]", values)
	}
}

func TestCounterRates(t *testing.T) {
	rates := counterRates([]metrics.Sample{
		{Timestamp: 0, Value: 100},
		{Timestamp: 10, Value: 130},
		{Timestamp: 20, Value: 190},
	})

	if len(rates) != 2 {
		t.Fatalf("got %d rates, want 2", len(rates))
	}
	if rates[0] != 3 || rates[1] != 6 {
		t.Fatalf("got rates %v, want [3 6]", rates)
	}
}

func TestAverageMinMax(t *testing.T) {
	values := []float64{40, 60, 80}

	if average(values) != 60 {
		t.Fatalf("got average %.2f, want 60.00", average(values))
	}
	if min(values) != 40 {
		t.Fatalf("got min %.2f, want 40.00", min(values))
	}
	if max(values) != 80 {
		t.Fatalf("got max %.2f, want 80.00", max(values))
	}
}

func TestRateUnit(t *testing.T) {
	if rateUnit("requests") != "requests/sec" {
		t.Fatalf("got rate unit %q, want %q", rateUnit("requests"), "requests/sec")
	}
}
