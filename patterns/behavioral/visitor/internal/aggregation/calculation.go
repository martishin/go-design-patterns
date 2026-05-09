package aggregation

import "github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"

func sampleValues(samples []metrics.Sample) []float64 {
	values := make([]float64, 0, len(samples))
	for _, sample := range samples {
		values = append(values, sample.Value)
	}

	return values
}

func counterRates(samples []metrics.Sample) []float64 {
	rates := make([]float64, 0, len(samples)-1)
	for index := 1; index < len(samples); index++ {
		previous := samples[index-1]
		current := samples[index]

		rates = append(rates, (current.Value-previous.Value)/float64(current.Timestamp-previous.Timestamp))
	}

	return rates
}

func average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	total := 0.0
	for _, value := range values {
		total += value
	}

	return total / float64(len(values))
}

func min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	minimum := values[0]
	for _, value := range values[1:] {
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

func max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	maximum := values[0]
	for _, value := range values[1:] {
		if value > maximum {
			maximum = value
		}
	}

	return maximum
}

func rateUnit(unit string) string {
	return unit + "/sec"
}

func copyResults(results []metrics.Result) []metrics.Result {
	return append([]metrics.Result(nil), results...)
}
