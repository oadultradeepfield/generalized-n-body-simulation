package simulation

import (
	"fmt"

	"github.com/oadultradeepfield/n-body-orbit-simulation/go/internal/body"
)

func InitializeSystem(cfg *Config) ([]*body.Body, error) {
	var system []*body.Body

	for _, b := range cfg.Bodies {
		var pos, vel []float64
		var err error

		switch cfg.CoordinatesType {
		case "spherical":
			pos, vel = ConvertSphericalToCartesian(b.Position, b.Velocity)
		case "cartesian":
			pos = make([]float64, 3)
			vel = make([]float64, 3)
			copy(pos, b.Position)
			copy(vel, b.Velocity)
		default:
			return nil, fmt.Errorf("invalid coordinate system: %s", cfg.CoordinatesType)
		}

		body, err := body.NewBody(b.Mass, pos, vel)
		if err != nil {
			return nil, fmt.Errorf("error creating body: %w", err)
		}
		system = append(system, body)
	}

	return system, nil
}
