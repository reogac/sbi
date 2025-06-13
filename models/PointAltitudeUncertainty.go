/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointAltitudeUncertainty struct {
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Point               GeographicalCoordinates `json:"point"`
	Shape               SupportedGADShapes      `json:"shape"`
	Altitude            float64                 `json:"altitude"`
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
	Confidence          int                     `json:"confidence"`
}
