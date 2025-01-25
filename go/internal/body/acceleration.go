package body

import "math"

func (b *Body) Acceleration(bodies []*Body, G float64) []float64 {
	acceleration := make([]float64, 3)
	const softening = 1e-8

	for _, other := range bodies {
		if other == b {
			continue
		}

		dx := other.position[0] - b.position[0]
		dy := other.position[1] - b.position[1]
		dz := other.position[2] - b.position[2]

		distanceSquared := dx*dx + dy*dy + dz*dz
		distance := math.Sqrt(distanceSquared)
		force := G * b.mass * other.mass / distanceSquared

		acceleration[0] += force * dx / (distance * b.mass)
		acceleration[1] += force * dy / (distance * b.mass)
		acceleration[2] += force * dz / (distance * b.mass)
	}

	return acceleration
}
