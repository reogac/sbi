/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PointAltitudeUncertainty struct {
	Shape               SupportedGADShapes      `json:"shape"`
	UncertaintyAltitude float64                 `json:"uncertaintyAltitude"`
	Confidence          int                     `json:"confidence"`
	VConfidence         *int                    `json:"vConfidence,omitempty"`
	Point               GeographicalCoordinates `json:"point"`
	Altitude            float64                 `json:"altitude"`
	UncertaintyEllipse  UncertaintyEllipse      `json:"uncertaintyEllipse"`
}
