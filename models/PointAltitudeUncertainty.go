/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointAltitudeUncertainty struct {
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
	Confidence          int                     `json:"confidence"`
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Shape               SupportedGADShapes      `json:"shape"`
	Point               GeographicalCoordinates `json:"point"`
	Altitude            float64                 `json:"altitude"`
}
