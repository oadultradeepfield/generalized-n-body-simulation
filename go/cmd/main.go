package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oadultradeepfield/n-body-orbit-simulation/go/internal/simulation"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <config_file.json>")
	}

	configFile := os.Args[1]

	cfg, err := simulation.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	bodySystem, err := simulation.InitializeSystem(cfg)
	if err != nil {
		log.Fatalf("Error initializing system: %v", err)
	}

	fmt.Printf("Simulation started with %d bodies...\n", len(bodySystem))

	startTime := time.Now()

	simulation.RunSimulation(bodySystem, cfg.Configuration.Dt, cfg.Configuration.TotalTime,
		cfg.Configuration.OutputFile, cfg.Configuration.G, cfg.Configuration.CollisionDistance)

	duration := time.Since(startTime)

	fmt.Printf("Simulation completed. Results saved to %s\n", cfg.Configuration.OutputFile)
	fmt.Printf("Time taken: %s\n", duration.Round(time.Millisecond))
}
