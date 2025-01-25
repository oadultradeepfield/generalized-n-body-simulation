package body

import "math"

func (b *Body) Collision(bodies []*Body, collisionDistance float64) {
	for _, other := range bodies {
		if other == b {
			continue
		}

		dx := other.position[0] - b.position[0]
		dy := other.position[1] - b.position[1]
		dz := other.position[2] - b.position[2]

		distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

		if distance < collisionDistance {
			for i := 0; i < 3; i++ {
				v1 := b.velocity[i]
				v2 := other.velocity[i]
				m1 := b.mass
				m2 := other.mass

				newV1 := ((m1-m2)*v1 + 2*m2*v2) / (m1 + m2)
				newV2 := ((m2-m1)*v2 + 2*m1*v1) / (m1 + m2)

				b.velocity[i] = newV1
				other.velocity[i] = newV2
			}
		}
	}
}
