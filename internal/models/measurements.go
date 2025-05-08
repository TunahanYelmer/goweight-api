package models

import "time"

// WeightReading represents a full data snapshot from the scale.
type WeightReading struct {
	Weight       float64   `json:"weight"`                  // in kg
	BodyFat      *float64  `json:"body_fat,omitempty"`      // in %
	MuscleMass   *float64  `json:"muscle_mass,omitempty"`   // in kg
	BoneMass     *float64  `json:"bone_mass,omitempty"`     // in kg
	WaterPercent *float64  `json:"water_percent,omitempty"` // in %
	BMI          *float64  `json:"bmi,omitempty"`           // Body Mass Index
	BMR          *float64  `json:"bmr,omitempty"`           // Basal Metabolic Rate
	Unit         string    `json:"unit"`                    // Unit of weight (e.g., kg, lb)
	Timestamp    time.Time `json:"timestamp"`               // Measurement timestamp
}
