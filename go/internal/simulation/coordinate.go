package simulation

import "math"

func ConvertSphericalToCartesian(sphericalPos, sphericalVel []float64) ([]float64, []float64) {
	r := sphericalPos[0]
	theta := sphericalPos[1]
	phi := sphericalPos[2]

	vr := sphericalVel[0]
	vtheta := sphericalVel[1]
	vphi := sphericalVel[2]

	pos := []float64{
		r * math.Cos(theta) * math.Sin(phi),
		r * math.Sin(theta) * math.Sin(phi),
		r * math.Cos(phi),
	}

	vel := []float64{
		vr*math.Sin(phi)*math.Cos(theta) - vtheta*math.Sin(theta) + vphi*math.Cos(theta)*math.Cos(phi),
		vr*math.Sin(phi)*math.Sin(theta) + vtheta*math.Cos(theta) + vphi*math.Cos(theta)*math.Sin(phi),
		vr*math.Cos(phi) - vphi*math.Sin(phi),
	}

	return pos, vel
}
