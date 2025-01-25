package body

import (
	"errors"
)

type Body struct {
	mass     float64
	position []float64
	velocity []float64
}

func NewBody(mass float64, position, velocity []float64) (*Body, error) {
	if mass <= 0 {
		return nil, errors.New("Mass must be a positive value")
	}

	return &Body{
		mass:     mass,
		position: position,
		velocity: velocity,
	}, nil
}

func (b *Body) Mass() float64       { return b.mass }
func (b *Body) Position() []float64 { return b.position }
func (b *Body) Velocity() []float64 { return b.velocity }
