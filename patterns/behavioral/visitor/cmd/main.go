package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/internal/aggregation"
	"github.com/martishin/go-design-patterns/patterns/behavioral/visitor/pkg/metrics"
)

type resultVisitor interface {
	metrics.MetricVisitor
	Results() []metrics.Result
}

func main() {
	monitoredMetrics := []metrics.Metric{
		mustMetric(metrics.NewGaugeMetric("cpu_usage", "percent", []metrics.Sample{
			{Timestamp: 0, Value: 40},
			{Timestamp: 10, Value: 60},
			{Timestamp: 20, Value: 80},
		})),
		mustMetric(metrics.NewCounterMetric("http_requests_total", "requests", []metrics.Sample{
			{Timestamp: 0, Value: 100},
			{Timestamp: 10, Value: 130},
			{Timestamp: 20, Value: 190},
		})),
	}

	printTimeSeries(monitoredMetrics)
	fmt.Println()
	printResults("Average visitor", aggregation.NewAverageVisitor(), monitoredMetrics)
	fmt.Println()
	printResults("Min visitor", aggregation.NewMinVisitor(), monitoredMetrics)
	fmt.Println()
	printResults("Max visitor", aggregation.NewMaxVisitor(), monitoredMetrics)
}

func mustMetric(metric metrics.Metric, err error) metrics.Metric {
	if err != nil {
		log.Fatal(err)
	}

	return metric
}

func printTimeSeries(monitoredMetrics []metrics.Metric) {
	fmt.Println("Time series:")
	for _, metric := range monitoredMetrics {
		fmt.Printf("- %s %s (%s):", metric.Name(), metric.Unit(), metricKind(metric))
		for _, sample := range metric.Samples() {
			fmt.Printf(" [%d]=%.2f", sample.Timestamp, sample.Value)
		}
		fmt.Println()

		counter, ok := metric.(*metrics.CounterMetric)
		if ok {
			printCounterRates(counter)
		}
	}
}

func printCounterRates(counter *metrics.CounterMetric) {
	samples := counter.Samples()

	fmt.Print("  rates:")
	for index := 1; index < len(samples); index++ {
		previous := samples[index-1]
		current := samples[index]
		rate := (current.Value - previous.Value) / float64(current.Timestamp-previous.Timestamp)

		fmt.Printf(" [%d-%d]=%.2f %s/sec", previous.Timestamp, current.Timestamp, rate, counter.Unit())
	}
	fmt.Println()
}

func metricKind(metric metrics.Metric) string {
	switch metric.(type) {
	case *metrics.GaugeMetric:
		return "gauge"
	case *metrics.CounterMetric:
		return "counter"
	default:
		return "metric"
	}
}

func printResults(title string, visitor resultVisitor, monitoredMetrics []metrics.Metric) {
	for _, metric := range monitoredMetrics {
		metric.Accept(visitor)
	}

	fmt.Printf("%s:\n", title)
	for _, result := range visitor.Results() {
		fmt.Printf("- %s %s: %.2f %s\n", result.Metric, result.Operation, result.Value, result.Unit)
	}
}
