package state

import (
	"time"
)

type Temperature struct {
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

// Struktura stanu systemu
type State struct {
	Id      string `json:"id" rethinkdb:"id,omitempty"`
	Heating struct {
		Temperature  Temperature `json:"temperature"`
		HeatingCurve Temperature `json:"heating_curve"`
	} `json:"heating"`
	Ventilation struct {
		ExtractionTemperature Temperature `json:"extraction_temperature"`
		SupplyTemperature     Temperature `json:"supply_temperature"`
	} `json:"ventilation"`
	Timestamp time.Time `json:"timestamp"`
}
