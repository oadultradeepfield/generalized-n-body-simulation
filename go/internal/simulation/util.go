package simulation

import (
	"fmt"
	"log"
	"os"

	"github.com/oadultradeepfield/n-body-orbit-simulation/go/internal/body"
)

func PrintState(bodies []*body.Body, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file %s: %v", filename, err)
		return
	}
	defer file.Close()

	for _, body := range bodies {
		pos := body.Position()
		_, err := fmt.Fprintf(file, "%f, %f, %f\n", pos[0], pos[1], pos[2])
		if err != nil {
			log.Printf("Error writing to file %s: %v", filename, err)
			return
		}
	}
}

func RunSimulation(bodies []*body.Body, dt, totalTime float64, filename string, G, collisionDistance float64) {
	steps := int(totalTime / dt)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Error clearing output file: %v", err)
		return
	}
	file.Close()

	for i := 0; i < steps; i++ {
		RungeKuttaStep(bodies, dt, G, collisionDistance)
		PrintState(bodies, filename)
	}
}

func ArrayToFixed(s []float64) [3]float64 {
	return [3]float64{s[0], s[1], s[2]}
}
