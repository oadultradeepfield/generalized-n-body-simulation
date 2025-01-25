package simulation

import "github.com/oadultradeepfield/n-body-orbit-simulation/go/internal/body"

func RungeKuttaStep(bodies []*body.Body, dt, G, collisionDistance float64) {
	n := len(bodies)

	k1Pos := make([][3]float64, n)
	k1Vel := make([][3]float64, n)
	k2Pos := make([][3]float64, n)
	k2Vel := make([][3]float64, n)
	k3Pos := make([][3]float64, n)
	k3Vel := make([][3]float64, n)
	k4Pos := make([][3]float64, n)
	k4Vel := make([][3]float64, n)

	copyBodies := func() []*body.Body {
		tmp := make([]*body.Body, n)
		for i, b := range bodies {
			pos := make([]float64, 3)
			copy(pos, b.Position())
			vel := make([]float64, 3)
			copy(vel, b.Velocity())
			tmp[i], _ = body.NewBody(b.Mass(), pos, vel)
		}
		return tmp
	}

	for i, b := range bodies {
		copy(k1Pos[i][:], b.Velocity())
		k1Vel[i] = ArrayToFixed(b.Acceleration(bodies, G))
	}

	tmp := copyBodies()
	for i := range tmp {
		for j := 0; j < 3; j++ {
			tmp[i].Position()[j] += 0.5 * dt * k1Pos[i][j]
			tmp[i].Velocity()[j] += 0.5 * dt * k1Vel[i][j]
		}
	}

	for i, b := range tmp {
		copy(k2Pos[i][:], b.Velocity())
		k2Vel[i] = ArrayToFixed(b.Acceleration(tmp, G))
	}

	tmp = copyBodies()
	for i := range tmp {
		for j := 0; j < 3; j++ {
			tmp[i].Position()[j] += 0.5 * dt * k2Pos[i][j]
			tmp[i].Velocity()[j] += 0.5 * dt * k2Vel[i][j]
		}
	}
	for i, b := range tmp {
		copy(k3Pos[i][:], b.Velocity())
		k3Vel[i] = ArrayToFixed(b.Acceleration(tmp, G))
	}

	tmp = copyBodies()
	for i := range tmp {
		for j := 0; j < 3; j++ {
			tmp[i].Position()[j] += dt * k3Pos[i][j]
			tmp[i].Velocity()[j] += dt * k3Vel[i][j]
		}
	}
	for i, b := range tmp {
		copy(k4Pos[i][:], b.Velocity())
		k4Vel[i] = ArrayToFixed(b.Acceleration(tmp, G))
	}

	for i, b := range bodies {
		for j := 0; j < 3; j++ {
			b.Position()[j] += dt / 6.0 * (k1Pos[i][j] + 2*k2Pos[i][j] + 2*k3Pos[i][j] + k4Pos[i][j])
			b.Velocity()[j] += dt / 6.0 * (k1Vel[i][j] + 2*k2Vel[i][j] + 2*k3Vel[i][j] + k4Vel[i][j])
		}
	}

	for _, b := range bodies {
		b.Collision(bodies, collisionDistance)
	}
}
