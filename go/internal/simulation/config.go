package simulation

import (
	"encoding/json"
	"os"
)

type Config struct {
	Configuration   ConfigInfo  `json:"config"`
	CoordinatesType string      `json:"coordinates_type"`
	Bodies          []BodyInput `json:"bodies"`
}

type ConfigInfo struct {
	G                 float64 `json:"G"`
	Dt                float64 `json:"dt"`
	TotalTime         float64 `json:"total_time"`
	OutputFile        string  `json:"filename"`
	CollisionDistance float64 `json:"collision_distance"`
}

type BodyInput struct {
	Name     string    `json:"_name"`
	Mass     float64   `json:"mass"`
	Position []float64 `json:"position"`
	Velocity []float64 `json:"velocity"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
