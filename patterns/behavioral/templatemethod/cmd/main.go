package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/internal/importer"
	"github.com/martishin/go-design-patterns/patterns/behavioral/templatemethod/pkg/document"
)

func main() {
	csvImporter, err := importer.NewCSVImporter("customers.csv", "id,name\n1,Alice\n2,Bob")
	if err != nil {
		log.Fatal(err)
	}
	runImport(csvImporter)

	fmt.Println()

	jsonImporter, err := importer.NewJSONImporter(
		"customers.json",
		`[{"id":"3","name":"Carol"},{"id":"4","name":"Dan"}]`,
	)
	if err != nil {
		log.Fatal(err)
	}
	runImport(jsonImporter)
}

func runImport(formatImporter document.Importer) {
	pipeline, err := importer.NewPipeline(formatImporter)
	if err != nil {
		log.Fatal(err)
	}

	report, err := pipeline.Import()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s import:\n", report.Format)
	for _, step := range report.Steps {
		fmt.Printf("- %s\n", step)
	}
	fmt.Printf("Imported records: %v\n", report.Records)
}
