package main

import (
	"fmt"
	"log"

	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/internal/monitor"
	"github.com/martishin/go-design-patterns/patterns/behavioral/observer/internal/subscriber"
)

func main() {
	fileMonitor := monitor.NewFileMonitor()
	indexer := subscriber.NewIndexer("indexer")
	backup := subscriber.NewBackupService("backup")

	must(fileMonitor.Subscribe(indexer))
	must(fileMonitor.Subscribe(backup))

	fmt.Println("Writing file: app/config.yaml")
	must(fileMonitor.Write("app/config.yaml"))
	printEntries(indexer.Updates())
	printEntries(backup.Backups())

	fmt.Println()
	fmt.Println("Backup service unsubscribed")
	fileMonitor.Unsubscribe(backup)

	fmt.Println("Deleting file: app/config.yaml")
	must(fileMonitor.Delete("app/config.yaml"))
	printLatest(indexer.Updates())
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printEntries(entries []string) {
	for _, entry := range entries {
		fmt.Printf("- %s\n", entry)
	}
}

func printLatest(entries []string) {
	if len(entries) == 0 {
		return
	}

	fmt.Printf("- %s\n", entries[len(entries)-1])
}
